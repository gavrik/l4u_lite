package main

import (
	"database/sql"
	"errors"
	"fmt"
	"lib"
)

// ErrAdminToken - Error
var ErrAdminToken = errors.New("Can't populate Admin Tokens Map")

// APISQLiteDB -
type APISQLiteDB struct {
	lib.SQLiteDBImplementation
}

// GetAdminTokens -
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

// PutLink -
func (impl *APISQLiteDB) PutLink(newLink *Link) error {
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

// GetLink -
func (impl *APISQLiteDB) GetLink(linkHash string, link *Link) error {
	rows, err := impl.Db.Query("select short_link, long_link, is_enabled, created_on from default_links where short_link = $1", linkHash)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&link.ShortLink, &link.LongLink, &link.IsEnabled, &link.CreationOn)
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
