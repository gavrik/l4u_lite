package lib

// SQLite3 implementation
import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	sqlite3 "github.com/mattn/go-sqlite3"
)

// ErrCallDB - Error
var ErrCallDB = errors.New("Can't execute select statement")

// ErrLinkNotFound - Error
var ErrLinkNotFound = errors.New("Link not found")

// ErrWrongDBStructure - Error
var ErrWrongDBStructure = errors.New("Wrong database structure")

// ErrUndefined - Error
var ErrUndefined = errors.New("Strange Error")

// ErrWrongDBVersion - Error
var ErrWrongDBVersion = errors.New("Wrong database version")

// ErrROOTtoken - Error
var ErrROOTtoken = errors.New("Can't generate root token")

// ErrInsert - Error
var ErrInsert = errors.New("Can't insert values")

// SQLiteDB - highlevel DB interface
type SQLiteDB interface {
	SQLInit() error
	Close()
	CheckDBversion() (int, error)
	GetLongLink(shortLink string, domain string, longLink *LongLink) error
	CreateROOToken() (string, error)
}

// LongLink - Error
type LongLink struct {
	id       int
	LongLink string
}

func init() {
	v, n, s := sqlite3.Version()
	fmt.Printf("SQLite version information: %s %d %s \n", v, n, s)
}

// SQLiteDBImplementation - SQLite data access implementation
type SQLiteDBImplementation struct {
	Db         *sql.DB
	ConnString string
	IsWritable bool
}

// SQLInit - Init database instance
func (impl *SQLiteDBImplementation) SQLInit() error {
	var err error
	impl.Db, err = sql.Open("sqlite3", impl.ConnString)
	return err
}

// Close - Close database instance
func (impl *SQLiteDBImplementation) Close() {
	impl.Db.Close()
}

// CheckDBversion - Check database schema version
func (impl *SQLiteDBImplementation) CheckDBversion() (int, error) {
	var row string
	rows, err := impl.Db.Query("select val from settings where key = 'VERSION' ")
	defer rows.Close()
	if err != nil {
		return 0, ErrCallDB
	}
	for rows.Next() {
		err = rows.Scan(&row)
		if err != nil {
			return 0, ErrWrongDBStructure
		}
	}

	return strconv.Atoi(row)
}

// GetLongLink - Get Long link from db
func (impl *SQLiteDBImplementation) GetLongLink(shortLink string, domain string, longLink *LongLink) error {
	var err error
	var rows *sql.Rows
	if domain == "" {
		rows, err = impl.Db.Query("select id, long_link from default_links where short_link = $1 and domain_id is null and is_enabled = 1", shortLink)
	} else {
		rows, err = impl.Db.Query("select id, long_link from default_links where short_link = $1 and domain_id = $2 and is_enabled = 1", shortLink, domain)
	}
	defer rows.Close()
	if err != nil {
		return ErrCallDB
	}
	for rows.Next() {
		err = rows.Scan(longLink.id, longLink.LongLink)
		if err != nil {
			return ErrLinkNotFound
		}
	}
	return nil
}

// CreateROOToken - Create root token for API
func (impl *SQLiteDBImplementation) CreateROOToken() (string, error) {
	token, err := GetUUID()
	if err != nil {
		return "", ErrROOTtoken
	}
	_, err = impl.Db.Query("insert into admin_tokens (token, token_description, is_Root) values( $1, 'ROOT token', 1)", token)
	if err != nil {
		return "", ErrInsert
	}
	return token, nil
}

// NewDB - Create DB instance
func newDB(dbPath string, dbMode string) SQLiteDB {
	db := new(SQLiteDBImplementation)
	db.ConnString = fmt.Sprintf("%s?mode=%s", dbPath, dbMode)
	db.IsWritable = false
	return db
}

// OpenDB - Open new SQLite3 read only database
func OpenDB(dbPath string) SQLiteDB {
	db := newDB(dbPath, "ro")
	db.SQLInit()
	return db
}

// OpenDBrw - Open new SQLite3 database for read and write operations
func OpenDBrw(dbPath string) SQLiteDB {
	db := newDB(dbPath, "rw")
	db.SQLInit()
	return db
}

// ChackDBVersion - Check database version
func ChackDBVersion(dbPath string) (int, error) {
	db := OpenDB(dbPath)
	dbversion, err := db.CheckDBversion()
	if err != nil {
		return 0, err
	}
	if dbversion != DatabaseVersion {
		return dbversion, ErrWrongDBVersion
	}
	db.Close()
	return dbversion, nil
}
