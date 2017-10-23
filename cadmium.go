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

package main

import (
	// local
	"github.com/pztrn/cadmium/configuration"
	"github.com/pztrn/cadmium/context"
	"github.com/pztrn/cadmium/database"
	"github.com/pztrn/cadmium/eventer"
	"github.com/pztrn/cadmium/ui"
)

func main() {
	c := context.New()
	c.Initialize()

	c.Log.Infoln("This is Cadmium, version 0.0.1")

	configuration.New(c)
	database.New(c)
	eventer.New(c)
	c.Config.LaterLoadConfig()

	// Launch user interface.
	ui.New(c)
}
