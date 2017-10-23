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

var start_schema = `
DROP TABLE IF EXISTS database;
CREATE TABLE database (
    version         VARCHAR(10)     NOT NULL
);

DROP TABLE IF EXISTS configuration;
CREATE TABLE configuration (
    key             VARCHAR(1024)    NOT NULL,
    value           VARCHAR(8192)    NOT NULL
);

INSERT INTO database (version) VALUES (1);
`

// Migrate database to latest version.
// ToDo: make it more good :).
func migrate_full(db *Database, version int) {
	c.Log.Infoln("Starting database migrations...")
	if version < 1 {
		start_to_one(db)
		version = 1
	}
	/*if version == 1 {
		one_to_two(db)
		version = 2
	}*/
}

// Initial database structure.
func start_to_one(db *Database) {
	c.Log.Infoln("Upgrading database from 0 to 1...")
	db.Db.MustExec(start_schema)
}
