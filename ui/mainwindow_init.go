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

	// local
	"github.com/pztrn/cadmium/eventer/event"

	// other
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/webengine"
	"github.com/therecipe/qt/widgets"
)

const (
	DEFAULT_WIN_POS_X = 50
	DEFAULT_WIN_POS_Y = 50

	DEFAULT_WIN_WIDTH  = 1000
	DEFAULT_WIN_HEIGHT = 600
)

// Initialize initializes main window.
func (m *MainWindow) Initialize() {
	c.Log.Infoln("Initializing main window...")

	var err error
	m.app = widgets.NewQApplication(len(os.Args), os.Args)
	m.app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	m.app.SetAttribute(core.Qt__AA_UseHighDpiPixmaps, true)
	core.QCoreApplication_SetOrganizationDomain("name")
	core.QCoreApplication_SetOrganizationName("pztrn")
	core.QCoreApplication_SetApplicationName("cadmium")
	core.QCoreApplication_SetApplicationVersion("0.0.1")

	m.window = widgets.NewQMainWindow(nil, 0)
	m.window.SetWindowTitle("Cadmium IM")

	// Restoring window position.
	savedWinPosXStr := c.Config.GetValue("/mainwindow/position_x")
	savedWinPosYStr := c.Config.GetValue("/mainwindow/position_y")
	winPosX, err := strconv.Atoi(savedWinPosXStr)
	if err != nil {
		winPosX = DEFAULT_WIN_POS_X
	}
	winPosY, err := strconv.Atoi(savedWinPosYStr)
	if err != nil {
		winPosY = DEFAULT_WIN_POS_Y
	}

	// Restoring window size.
	savedWinSizeWidthStr := c.Config.GetValue("/mainwindow/width")
	savedWinSizeHeightStr := c.Config.GetValue("/mainwindow/height")
	m.windowWidth, err = strconv.Atoi(savedWinSizeWidthStr)
	if err != nil {
		m.windowWidth = DEFAULT_WIN_WIDTH
	}
	m.windowHeight, err = strconv.Atoi(savedWinSizeHeightStr)
	if err != nil {
		m.windowHeight = DEFAULT_WIN_HEIGHT
	}

	m.window.SetGeometry2(winPosX, winPosY, m.windowWidth, m.windowHeight)

	widget := widgets.NewQWidget(nil, 0)
	loader := uitools.NewQUiLoader(nil)
	mwfile := core.NewQFile2(":/qml/ui/mainwindow.ui")

	mwfile.Open(core.QIODevice__ReadOnly)
	mainWidget := loader.Load(mwfile, widget)
	mwfile.Close()
	m.window.SetCentralWidget(mainWidget)

	// Resize event handler.
	m.window.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		m.windowResizeHandler(event)
	})

	// Window move event handler.
	m.window.ConnectMoveEvent(func(event *gui.QMoveEvent) {
		m.windowMoveHandler(event)
	})

	// Add events.
	m.initializeCoreEvents()

	// Menu.
	m.initializeMenu()

	// Chat history things.
	chatHistoryLayout := widgets.NewQVBoxLayout()
	chatHistoryLayout.SetContentsMargins(0, 0, 0, 0)
	chatHistoryWidget := widgets.NewQWidgetFromPointer(mainWidget.FindChild("ChatWidget", core.Qt__FindChildrenRecursively).Pointer())
	chatHistoryWidget.SetLayout(chatHistoryLayout)

	m.chatHistory = webengine.NewQWebEngineView(nil)
	chatHistoryLayout.AddWidget(m.chatHistory, 0, core.Qt__AlignVCenter&core.Qt__AlignHCenter)

	// Input field.
	m.messageInput = widgets.NewQTextEditFromPointer(mainWidget.FindChild("MessageInput", core.Qt__FindChildrenRecursively).Pointer())
	// Input field height.
	m.messageInputTextChanged()
	m.messageInput.ConnectTextChanged(m.messageInputTextChanged)

	m.window.Show()

	// Splitter between roster and chat.
	m.rosterAndChatSplitter = widgets.NewQSplitterFromPointer(mainWidget.FindChild("RosterAndChatSplitter", core.Qt__FindChildrenRecursively).Pointer())
	m.rosterAndChatSplitter.ConnectSplitterMoved(m.rosterAndChatSplitterMoved)
	// Restore it's position.
	// Warning: this should be done exactly here as it recalculates
	// sizes of widgets!
	rosterAndChatSplitterPosStr := c.Config.GetValue("/mainwindow/rosterandchatsplitterpos")
	if rosterAndChatSplitterPosStr != "" {
		rosterAndChatSplitterPos, _ := strconv.Atoi(rosterAndChatSplitterPosStr)
		m.rosterAndChatSplitterPos = rosterAndChatSplitterPos
	} else {
		g := m.window.Geometry()
		w := g.Width()
		m.rosterAndChatSplitterPos = w - (w - 250)
	}
	c.Log.Debugln("Roster and chats splitter position:", m.rosterAndChatSplitterPos)

	m.rosterAndChatSplitter.MoveSplitter(m.rosterAndChatSplitterPos, 1)

	widgets.QApplication_Exec()
}

func (m *MainWindow) initializeCoreEvents() {
	// Close event.
	coreCloseEvent := &event.Event{
		Name:        "closeCadmium",
		Description: "Final event handler for closeCadmium event",
		Handler:     m.closeAppEventHandler,
		EventID:     999,
	}
	c.Eventer.AddEventHandler("closeCadmium", coreCloseEvent)
	// Close button on window.
	m.window.ConnectCloseEvent(func(event *gui.QCloseEvent) {
		m.closeApp(true)
	})
}

func (m *MainWindow) initializeMenu() {
	c.Log.Debugln("Connecting menu items...")
	actionQuit := widgets.NewQActionFromPointer(m.window.CentralWidget().FindChild("action_Quit", core.Qt__FindChildrenRecursively).Pointer())
	actionQuit.ConnectTriggered(m.closeApp)

	actionAppLogs := widgets.NewQActionFromPointer(m.window.CentralWidget().FindChild("actionApplication_logs", core.Qt__FindChildrenRecursively).Pointer())
	actionAppLogs.ConnectTriggered(m.showApplicationLogsDialog)

	actionDebugInfo := widgets.NewQActionFromPointer(m.window.CentralWidget().FindChild("actionDebug_info", core.Qt__FindChildrenRecursively).Pointer())
	actionDebugInfo.ConnectTriggered(m.showDebugInfoDialog)
}
