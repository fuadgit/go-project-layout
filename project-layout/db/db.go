package db

import (
	"database/sql"

	"ibfd.org/app/cfg"
	log "ibfd.org/app/log4u"
	// TODO import database driver
	// e.g. _ "github.com/go-sql-driver/mysql" // mysql driver
)

// TODO Replace DB by database name
// e.g. If database name = repo_name. Replace DB by RepoMaster

// DB database. // TODO Replace DB by database name
type DB struct {
	db *sql.DB
}

// NewDB creates a DB handle. // TODO Replace NewDB name by databse name. e.g. NewRepoMaster
func NewDB(dbDef *cfg.DbDef) *DB {
	return &DB{openDatabase(dbDef)}
}

func openDatabase(dbDef *cfg.DbDef) *sql.DB {
	dbSrc := "???"                    // TODO add database connection string
	db, err := sql.Open("???", dbSrc) // TODO replace ??? by database name. e.g. mysql
	if err != nil {
		log.Fatalf("failed to open database %s:%d/%s: [%v]", dbDef.Host, dbDef.Port, dbDef.Database, err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to open database %s:%d/%s: [%v]", dbDef.Host, dbDef.Port, dbDef.Database, err)
	}
	return db
}

// Close closes all database connections
func (d *DB) Close() {
	d.db.Close()
}
