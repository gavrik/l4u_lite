package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"lib"
	"strings"
)

// ErrAdminToken - Error
var ErrAdminToken = errors.New("Can't populate Admin Tokens Map")

// ErrDbNotWritable - Error
var ErrDbNotWritable = errors.New("DB is not open for write operations")

// ErrLinkNotFound - Error
var ErrLinkNotFound = errors.New("Link hash was not found in database")

// ErrDBExists - Error
var ErrDBExists = errors.New("Database Exists. Do not need create fresh schema")

// APISQLiteDB -
type APISQLiteDB struct {
	lib.SQLiteDBImplementation
}

// CreateAdminToken - create tokens for REST access
func (impl *APISQLiteDB) CreateAdminToken(tokenHash string, isRoot bool, desc string, expireAt int) error {
	var numRows int64 = 0
	var err error
	var stm *sql.Stmt
	var result sql.Result
	stm, err = impl.Db.Prepare("insert into admin_tokens (token, token_description, is_Root, expire_at) values(?,?,?,?)")
	if err != nil {
		return err
	}
	defer stm.Close()
	result, err = stm.Exec(tokenHash, desc, isRoot, expireAt)
	if err != nil {
		return err
	}
	numRows, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if numRows > 0 {
		return nil
	}
	return nil
}

// GetAdminTokens - load admin tokens to HashMap
func (impl *APISQLiteDB) GetAdminTokens(tokenMap map[string]AdminToken) error {
	rows, err := impl.Db.Query("select t.token, t.token_description, ifnull(d.domain_name,''), t.expire_at, t.is_Root from admin_tokens t left join domain d on t.domain_id = d.id;")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var token AdminToken
		err = rows.Scan(&token.Token, &token.TokenDescription, &token.Domain, &token.ExpireAt, &token.IsRoot)
		if err != nil {
			fmt.Println(err)
			return ErrAdminToken
		}
		fmt.Println(token)
		tokenMap[token.Token] = token
	}

	return err
}

// PutLink - Create new link record
func (impl *APISQLiteDB) PutLink(newLink *Link) error {
	if !impl.IsWritable {
		return ErrDbNotWritable
	}
	var rows *sql.Rows
	stm, err := impl.Db.Prepare("insert into default_links(short_link,long_link,is_enabled) values(?,?,?)")
	if err != nil {
		return err
	}
	_, err = stm.Exec(newLink.ShortLink, newLink.LongLink, newLink.IsEnabled)
	if err != nil {
		return err
	}
	rows, err = impl.Db.Query("select created_on from default_links where short_link = $1", newLink.ShortLink)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&newLink.CreationOn)
	}

	return err
}

// GetLink - Get full link info
func (impl *APISQLiteDB) GetLink(link *Link) error {
	rows, err := impl.Db.Query("select long_link, is_enabled, created_on from default_links where short_link = $1", link.ShortLink)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&link.LongLink, &link.IsEnabled, &link.CreationOn)
	}
	return err
}

// DeleteLink - Delete link record
func (impl *APISQLiteDB) DeleteLink(link *Link) error {
	if !impl.IsWritable {
		return ErrDbNotWritable
	}
	_, err := impl.Db.Exec("delete from default_links where short_link = $1", link.ShortLink)
	if err != nil {
		return err
	}
	return err
}

// UpdateLink - Update link parameters
func (impl *APISQLiteDB) UpdateLink(link *Link) error {
	if !impl.IsWritable {
		return ErrDbNotWritable
	}
	var n int64
	res, err := impl.Db.Exec("update default_links set long_link = $1, is_enabled = $2 where short_link = $3",
		link.LongLink, link.IsEnabled, link.ShortLink)
	if err != nil {
		return err
	}
	n, err = res.RowsAffected()
	if err != nil || n == 0 {
		err = ErrLinkNotFound
	}
	return err
}

func openDB(dbPath, mode string) *APISQLiteDB {
	db := new(APISQLiteDB)
	db.ConnString = fmt.Sprintf("%s?mode=%s", dbPath, mode)
	return db
}

// NewAPIDB - Create DB instance
func NewAPIDB(dbPath string) *APISQLiteDB {
	db := openDB(dbPath, "rw")
	db.IsWritable = true
	db.SQLInit()
	return db
}

// NewAPIDBro - Create read only DB instance
func NewAPIDBro(dbPath string) *APISQLiteDB {
	db := openDB(dbPath, "ro")
	db.IsWritable = false
	db.SQLInit()
	return db
}

// CreateFreshDB - create fresh database and load default structure to it
func CreateFreshDB(config *AppConfig) error {
	if !config.IsCreateDB {
		return nil
	}
	if lib.IsFileExixts(config.DatabasePath) {
		return ErrDBExists
	}

	DB := NewAPIDB(config.DatabasePath)
	defer DB.Db.Close()

	file, err := ioutil.ReadFile(lib.FreshDBSQLFile)
	if err != nil {
		fmt.Println(err)
	}
	queres := strings.Split(string(file), ";")

	for _, query := range queres {
		if config.IsDebug {
			fmt.Println("*** Execute sql statement")
			fmt.Println(query)
		}

		if len(query) > 4 {
			_, err := DB.Db.Exec(query)
			fmt.Println(err)
		}
	}
	token, err := lib.GetUUID()
	if err != nil {
		return err
	}
	err = DB.CreateAdminToken(token, true, "ROOT TOKEN. DO NOT DELETE IT!!!!", 0)
	if err != nil {
		return err
	}
	fmt.Printf("\n\n\n* * * * * INITIAL ROOT TOKEN * * * * *\n")
	fmt.Printf("\n %s \n\n", token)
	fmt.Printf("* * * * * * * * * * * * * * * * * * * *\n\n\n")

	return nil
}
