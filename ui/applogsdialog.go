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
	// other
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

func (m *MainWindow) showApplicationLogsDialog(checked bool) {
	c.Log.Debugln("Showing application logs...")

	widget := widgets.NewQWidget(nil, 0)
	loader := uitools.NewQUiLoader(nil)
	mwfile := core.NewQFile2(":/qml/ui/logsviewer.ui")

	mwfile.Open(core.QIODevice__ReadOnly)
	mainWidget := loader.Load(mwfile, widget)
	mwfile.Close()

	window := widgets.NewQDialog(nil, core.Qt__Dialog)
	window.SetWindowTitle("Logs viewer")
	window.SetMinimumHeight(700)
	window.SetMinimumWidth(650)
	mainvbox := widgets.NewQVBoxLayout()
	mainvbox.SetContentsMargins(0, 0, 0, 0)
	mainvbox.AddWidget(mainWidget, 0, core.Qt__AlignHCenter&core.Qt__AlignTop)
	window.SetLayout(mainvbox)

	window.Show()
}
