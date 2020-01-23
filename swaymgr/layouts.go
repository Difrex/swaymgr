package main

import (
	"github.com/Difrex/gosway/ipc"
)

// Layout is an
type Layout interface {
	// OnNew must receive an *ipc.Event
	// and do the container manipulation
	OnNew(*ipc.Event) error
	// OnFocus process the focus event
	OnFocus(*ipc.Event) error
	// Manage must store WorkspaceConfig in the database with
	// the workspace name, layout name and with the Managed: true
	Manage() error
}

// NewLayouts initilizes all the layouts
func NewLayouts(conn *ipc.SwayConnection, store *store) map[string]Layout {
	layouts := make(map[string]Layout)

	spiral := Layout(NewSpiralLayout(conn, store))
	left := Layout(NewLeftLayout(conn, store))
	layouts["spiral"] = spiral
	layouts["left"] = left

	return layouts
}

// Unmanage makes the currently focused workspace non-management
func (m *manager) Unmanage() error {
	ws, err := m.commandConn.GetFocusedWorkspace()
	if err != nil {
		return err
	}

	wc := WorkspaceConfig{
		Name:    ws.Name,
		Layout:  "",
		Managed: false,
	}

	if err := m.store.put([]byte(ws.Name), wc); err != nil {
		return err
	}

	return nil
}

type FiberLayout struct{}
type TopLayout struct{}
type BottomLayout struct{}
type RightLayout struct{}
