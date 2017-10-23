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

import (
	// stdlib
	"os"
	"path"
	"runtime"
	"strconv"

	// local
	"github.com/pztrn/cadmium/database/models"
	"github.com/pztrn/cadmium/eventer/event"
)

// Config provides access to configuration storage.
// It also stores some basic variables that defined in temporary
// configuration storage. They can be accessed with these keys:
//
//   * CONFIGPATH - application's configuration path, defaulting to
//     /Users/<user>/Library/Application Support/Cadmium os macOS,
//     ~/.config/Cadmium on Nixes and APPDATA/Roaming/Cadmiun on
//     Windows.
//     This path is used to build other paths, like database path,
//     caches path, etc.
//   * CACHEPATH - application's caches path.
type Config struct {
	// Persistent configuration storage that will be written in database.
	// Keys form: "module/parameter/..." e.g. "protocols/xmpp/conn1/jid".
	config map[string]string
	// Temporary configuration storage which will be lost on application
	// restart. Keys should be formatted same as in config map.
	tempconfig map[string]string
}

// Initialize initializes configuration storage and application paths
// for later access by e.g. Database.
func (cfg *Config) Initialize() {
	c.Log.Infoln("Initializing configuration...")

	cfg.config = make(map[string]string)
	cfg.tempconfig = make(map[string]string)

	// Figure out where to write our configuration, database and other
	// files.
	if runtime.GOOS == "darwin" {
		cfg.initializePathsMac()
	} else if runtime.GOOS == "windows" {
		cfg.initializePathsWin()
	} else if runtime.GOOS == "linux" {
		cfg.initializePathsNix()
	} else {
		c.Log.Fatalln("Cadmiun can be run only on Windows, Linux and macOS for now!")
	}

	c.Log.Infoln("Application's configuration and data path:", cfg.tempconfig["CONFIGPATH"])

	if _, err := os.Stat(cfg.tempconfig["CONFIGPATH"]); os.IsNotExist(err) {
		c.Log.Warnln("Path", cfg.tempconfig["CONFIGPATH"], "does not exists, creating...")
		os.MkdirAll(cfg.tempconfig["CONFIGPATH"], 0755)
		// As path does not exists - set tempconfig["FIRSTLAUNCH"].
		cfg.tempconfig["FIRSTLAUNCH"] = "1"
	}
}

// initializePathsMac initializes paths configuration as values in
// temporary configuration storage for macOS.
func (cfg *Config) initializePathsMac() {
	homePath := os.Getenv("HOME")

	cfg.tempconfig["CONFIGPATH"] = path.Join(homePath, "Library", "Application Support", "Cadmium")
	cfg.tempconfig["CACHEPATH"] = path.Join(cfg.tempconfig["CONFIGPATH"], "cache")
}

// initializePathsNix initializes paths configuration as values in
// temporary configuration storage for Nixes.
func (cfg *Config) initializePathsNix() {
	homePath := os.Getenv("HOME")

	cfg.tempconfig["CONFIGPATH"] = path.Join(homePath, ".config", "Cadmium")
	cfg.tempconfig["CACHEPATH"] = path.Join(cfg.tempconfig["CONFIGPATH"], "cache")
}

// initializePathsWin initializes paths configuration as values in
// temporary configuration storage for Windows.
func (cfg *Config) initializePathsWin() {
	homePathWithoutDrive := os.Getenv("HOMEPATH")
	homeDrive := os.Getenv("HOMEDRIVE")

	cfg.tempconfig["CONFIGPATH"] = path.Join(homeDrive, homePathWithoutDrive, "AppData", "Roaming", "Cadmium")
	cfg.tempconfig["CACHEPATH"] = path.Join(cfg.tempconfig["CONFIGPATH"], "cache")
}

// GetTempValue returns value for key from temporary configuration
// storage.
func (cfg *Config) GetTempValue(key string) string {
	value, found := cfg.tempconfig[key]
	if found {
		return value
	}

	return ""
}

// GetValue returns value for key from persistent configuration
// storage.
func (cfg *Config) GetValue(key string) string {
	value, found := cfg.config[key]
	if found {
		return value
	}

	return ""
}

// LaterLoadConfig loads configuration from already initialized database
// and adds some events handlers.
func (cfg *Config) LaterLoadConfig() {
	c.Log.Infoln("Loading configuration into memory...")

	// Add event for saving configuration to database.
	e := &event.Event{
		Name:        "saveConfiguration",
		Description: "Final event for saveConfiguration event",
		Handler:     cfg.saveConfiguration,
		EventID:     999,
	}
	c.Eventer.AddEventHandler("saveConfiguration", e)

	// Load configuration itself.
	db := c.Database.GetDatabase()
	var cfgs []models.Configuration
	err := db.Select(&cfgs, "SELECT * FROM configuration")
	if err != nil {
		c.Log.Errorln("Failed to load configuration:", err.Error())
		return
	}
	for i := range cfgs {
		cfg.config[cfgs[i].Key] = cfgs[i].Value
	}
	c.Log.Debugln("Loaded", strconv.Itoa(len(cfg.config)), "configuration parameters")
}

// Saves configuration to database.
func (cfg *Config) saveConfiguration(data map[string]string) {
	c.Log.Infoln("Saving configuration to database...")

	db := c.Database.GetDatabase()

	db.MustExec("DELETE FROM configuration")

	for k, v := range cfg.config {
		data := models.Configuration{
			Key:   k,
			Value: v,
		}
		db.NamedExec("INSERT INTO configuration (key, value) VALUES (:key, :value)", &data)
	}
}

// SetTempValue sets value for key in temporary configuration storage.
func (cfg *Config) SetTempValue(key string, value string) {
	cfg.tempconfig[key] = value
}

// SetValue sets value for key in persistent configuration storage.
func (cfg *Config) SetValue(key string, value string) {
	cfg.config[key] = value
}
