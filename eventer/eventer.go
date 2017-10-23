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
package eventer

import (
	// stdlib
	"fmt"
	"sync"

	// local
	"github.com/pztrn/cadmium/eventer/event"
)

type Eventer struct {
	// Events list.
	// Events will be executed in order of adding.
	// We expect maximum 999 events to be.
	events map[string]map[int]*event.Event
	// Mutex for events list.
	eventsMutex sync.Mutex
}

// AddEventHandler adds a handler for desired event. This handler will be
// appended to list of currently existing handlers. If there was no event
// with such name - list will be created and handler will be placed at
// position 0.
func (e *Eventer) AddEventHandler(event_name string, handler *event.Event) {
	c.Log.Debugln("Adding event handler", handler.Name, "for event", event_name)
	c.Log.Debugln("Handler:", fmt.Sprintf("%+v", handler))

	e.eventsMutex.Lock()
	_, event_found := e.events[event_name]
	if event_found {
		var found bool = false
		for i := range e.events[event_name] {
			if e.events[event_name][i].Name == handler.Name {
				found = true
			}
		}

		if found {
			c.Log.Warnln("Event", event_name, "already containing handler", handler.Name, ", you should remove it first. No action was taken.")
			e.eventsMutex.Unlock()
			return
		}
	} else {
		c.Log.Debugln("Initializing empty events chain for event", event_name)
		e.events[event_name] = make(map[int]*event.Event)
	}

	e.events[event_name][len(e.events[event_name])] = handler
	e.eventsMutex.Unlock()
}

// Initialize initializes Eventer.
func (e *Eventer) Initialize() {
	c.Log.Infoln("Initializing events handler...")
	e.events = make(map[string]map[int]*event.Event)
}

// LaunchEvent launches desired event.
func (e *Eventer) LaunchEvent(event_name string, data map[string]string) {
	e.eventsMutex.Lock()
	handlers, event_found := e.events[event_name]
	e.eventsMutex.Unlock()
	if event_found {
		c.Log.Debugln("Launching event", event_name)
		for i := 0; i <= 999; i++ {
			handler, found := handlers[i]
			if found {
				handler.Handler(data)
			}
		}
	} else {
		c.Log.Warnln("Event", event_name, "not found")
	}
}

// RemoveEventHandler removes event handler for desired event. If handler
// or event wasn't found - no action is taken.
func (e *Eventer) RemoveEventHandler(event_name string, handler_name string) {
	e.eventsMutex.Lock()
	handlers, event_found := e.events[event_name]
	if !event_found {
		c.Log.Warnln("Can't delete handler", handler_name, "for event", event_name, "- event wasn't found.")
		e.eventsMutex.Unlock()
		return
	}

	var event_id int = 0
	for i := range handlers {
		if handlers[i].Name == handler_name {
			c.Log.Debugln("Found event handler", handler_name, "at position", i)
			event_id = i
			break
		}
	}

	if event_id == 0 {
		c.Log.Warnln("Can't find ID for handler", handler_name, "registered for event", event_name)
		e.eventsMutex.Unlock()
		return
	}

	delete(e.events[event_name], event_id)
	for i := range handlers {
		// Maybe there is a better way than this?
		if i > event_id && event_id != 999 {
			e.events[event_name][i-1] = e.events[event_name][i]
			delete(e.events[event_name], i)
		}
	}
	e.eventsMutex.Unlock()
}
