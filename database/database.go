// Cadmium - multiprotocol crossplatform messenger.
// Copyright (c) 2017, Stanislav N aka pztrn.
// Copyright (c) 2017, Cadmium developers and contributors.
//
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package database

import (
	// stdlib
	"path"
	"runtime"
	"strconv"

	// local
	"github.com/pztrn/cadmium/database/models"

	// Other
	"github.com/jmoiron/sqlx"
	// We're using sqlite3 database.
	_ "github.com/mattn/go-sqlite3"
)

// Database gives access to database for whole application.
// It can be used to store configuration, history logs, etc.
type Database struct {
	// Opened database.
	Db *sqlx.DB
}

// GetDatabase returns pointer to database.
func (d *Database) GetDatabase() *sqlx.DB {
	return d.Db
}

// Initialize initializes database connection.
func (d *Database) Initialize() {
	c.Log.Infoln("Initializing Database...")

	runtime.LockOSThread()

	databasePath := path.Join(c.Config.GetTempValue("CONFIGPATH"), "database.sqlite3")
	db, err := sqlx.Connect("sqlite3", databasePath)
	if err != nil {
		c.Log.Fatalln(err.Error())
	}
	d.Db = db

	c.Log.Infoln("Database opened.")
}

// Migrate launches database migration to latest version.
func (d *Database) Migrate() {
	// Getting current database version.
	dbver := 0
	database := []models.Database{}
	d.Db.Select(&database, "SELECT * FROM database")
	if len(database) > 0 {
		c.Log.Debugln("Current database version:", database[0].Version)
		dbver, _ = strconv.Atoi(database[0].Version)
	} else {
		c.Log.Debugln("No database found, will create new one")
	}

	migrate_full(d, dbver)
}
