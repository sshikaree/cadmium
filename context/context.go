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
package context

import (
	// stdlib
	"os"

	// local
	"github.com/pztrn/cadmium/configuration/configurationinterface"
	"github.com/pztrn/cadmium/database/databaseinterface"
	"github.com/pztrn/cadmium/eventer/eventerinterface"

	// other
	"lab.pztrn.name/golibs/flagger"
	"lab.pztrn.name/golibs/mogrus"
)

// Context structure provides application-wide access to some things
// like Logger, CLI flags parser, configuration and more.
type Context struct {
	Config   configurationinterface.ConfigurationInterface
	Database databaseinterface.DatabaseInterface
	Eventer  eventerinterface.EventerInterface
	Flagger  *flagger.Flagger
	Log      *mogrus.LoggerHandler
}

// Initialize application-wide Context.
func (c *Context) Initialize() {
	l := mogrus.New()
	l.Initialize()
	c.Log = l.CreateLogger("opensaps")
	c.Log.CreateOutput("stdout", os.Stdout, true)

	c.Flagger = flagger.New(c.Log)
	c.Flagger.Initialize()
}

// RegisterDatabaseInterface registers database interface for using
// with other parts of Cadmium.
func (c *Context) RegisterDatabaseInterface(di databaseinterface.DatabaseInterface) {
	c.Database = di
	c.Database.Initialize()
	c.Database.Migrate()
}

// RegisterEventerInterface registers eventer interface for using
// with other parts of Cadmium.
func (c *Context) RegisterEventerInterface(ei eventerinterface.EventerInterface) {
	c.Eventer = ei
	c.Eventer.Initialize()
}

// RegisterConfigurationInterface registers configuration interface for
// using with other parts of Cadmium.
func (c *Context) RegisterConfigurationInterface(ci configurationinterface.ConfigurationInterface) {
	c.Config = ci
	c.Config.Initialize()
}
