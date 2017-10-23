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
package configuration

// ConfigHandler provides handler for ConfigurationInterface.
// It launching real functions from Config struct.
type ConfigHandler struct{}

func (ch ConfigHandler) Initialize() {
	config.Initialize()
}

func (ch ConfigHandler) GetTempValue(key string) string {
	return config.GetTempValue(key)
}

func (ch ConfigHandler) GetValue(key string) string {
	return config.GetValue(key)
}

func (ch ConfigHandler) LaterLoadConfig() {
	config.LaterLoadConfig()
}

func (ch ConfigHandler) SetTempValue(key string, value string) {
	config.SetTempValue(key, value)
}

func (ch ConfigHandler) SetValue(key string, value string) {
	config.SetValue(key, value)
}
