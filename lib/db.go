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

// SQLiteDB - highlevel DB interface
type SQLiteDB interface {
	SQLInit() error
	Close()
	CheckDBversion() (int, error)
	GetLongLink(shortLink string, domain string, longLink *LongLink) error
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
	db         *sql.DB
	ConnString string
}

// SQLInit - Init database instance
func (imp *SQLiteDBImplementation) SQLInit() error {
	var err error
	imp.db, err = sql.Open("sqlite3", imp.ConnString)
	return err
}

// Close - Close database instance
func (imp *SQLiteDBImplementation) Close() {
	imp.db.Close()
}

// CheckDBversion - Check database schema version
func (imp *SQLiteDBImplementation) CheckDBversion() (int, error) {
	var row string
	rows, err := imp.db.Query("select val from settings where key = 'VERSION' ")
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
func (imp *SQLiteDBImplementation) GetLongLink(shortLink string, domain string, longLink *LongLink) error {
	var err error
	var rows *sql.Rows
	if domain == "" {
		rows, err = imp.db.Query("select id, long_link from default_links where short_link = $1 and domain_id is null and is_enabled = 1", shortLink)
	} else {
		rows, err = imp.db.Query("select id, long_link from default_links where short_link = $1 and domain_id = $2 and is_enabled = 1", shortLink, domain)
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

// NewDB - Create DB instance
func NewDB(dbPath string) SQLiteDB {
	db := new(SQLiteDBImplementation)
	db.ConnString = fmt.Sprintf("%s?mode=ro", dbPath)
	return db
}

// OpenDB - Open new SQLite3 Database
func OpenDB(dbPath string) SQLiteDB {
	db := NewDB(dbPath)
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
		return 0, ErrWrongDBVersion
	}
	db.Close()
	return dbversion, nil
}
