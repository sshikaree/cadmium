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
package ui

import (
	// stdlib
	"os"
	"strconv"

	// other
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/webengine"
	"github.com/therecipe/qt/widgets"
)

// MainWindow represents main window.
type MainWindow struct {
	// The application.
	app *widgets.QApplication
	// Main window.
	window *widgets.QMainWindow
	// Vertical splitter between roster/search and chat things.
	rosterAndChatSplitter *widgets.QSplitter
	// Chat history.
	chatHistory *webengine.QWebEngineView
	// Message input field.
	messageInput *widgets.QTextEdit

	// Window-related sizes.
	windowWidth  int
	windowHeight int
	windowPosX   int
	windowPosY   int
	// Roster-chat splitter position.
	rosterAndChatSplitterPos int
}

// This function executed only if ALT+F4/CMD+Q/window close button
// was clicked.
func (m *MainWindow) closeApp(checked bool) {
	if checked {
		c.Log.Debugln("Closed from window button")
	}
	c.Log.Infoln("Closing Cadmium...")

	c.Eventer.LaunchEvent("closeCadmium", nil)
	os.Exit(0)
}

// This function launched at last when closeCadmium event is emitted.
// This function executes database update with configuration values
// and database dump to disk.
func (m *MainWindow) closeAppEventHandler(data map[string]string) {
	c.Log.Debugln("Launching final close event...")

	// Put window-related data to config.
	c.Config.SetValue("/mainwindow/position_x", strconv.Itoa(m.windowPosX))
	c.Config.SetValue("/mainwindow/position_y", strconv.Itoa(m.windowPosY))
	c.Config.SetValue("/mainwindow/width", strconv.Itoa(m.windowWidth))
	c.Config.SetValue("/mainwindow/height", strconv.Itoa(m.windowHeight))
	c.Config.SetValue("/mainwindow/rosterandchatsplitterpos", strconv.Itoa(m.rosterAndChatSplitterPos))

	c.Eventer.LaunchEvent("saveConfiguration", nil)
}

func (m *MainWindow) rosterAndChatSplitterMoved(pos int, index int) {
	c.Log.Debugln("Splitter moved! Pos:", pos, ", index:", index)
	m.rosterAndChatSplitterPos = pos

}

func (m *MainWindow) windowMoveHandler(event *gui.QMoveEvent) {
	m.windowPosY = event.Pos().Y()
	m.windowPosX = event.Pos().X()
	//c.Log.Debugln("MainWindow's position changed:", strconv.Itoa(m.windowPosX), strconv.Itoa(m.windowPosY))
}

func (m *MainWindow) windowResizeHandler(event *gui.QResizeEvent) {
	m.windowWidth = event.Size().Width()
	m.windowHeight = event.Size().Height()
	//c.Log.Debugln("MainWindow's WxH changed:", strconv.Itoa(m.windowWidth), strconv.Itoa(m.windowHeight))
}
