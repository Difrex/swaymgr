package main

import (
	"fmt"

	"github.com/Difrex/gosway/ipc"
)

type SpiralLayout struct {
	Conn  *ipc.SwayConnection
	store *store
}

// NewSpiralLayout initializes the SpiralLayout interface
func NewSpiralLayout(conn *ipc.SwayConnection, store *store) *SpiralLayout {
	layout := &SpiralLayout{}
	layout.Conn = conn
	layout.store = store
	return layout
}

// OnFocus ...
func (s *SpiralLayout) OnFocus(event *ipc.Event) error {
	return s.OnNew(event)
}

// OnNew places new container as a spiral
// it's calculates container width and heigh
// if the layout width is larger than height
// it splits container by horizontal side
// else splits it by vertical side
func (s *SpiralLayout) OnNew(event *ipc.Event) error {
	nodes, err := s.Conn.GetFocusedWorkspaceWindows()
	if err != nil {
		return err
	}
	var result ipc.Node
	for _, node := range nodes {
		if node.Focused {
			result = node
			break
		}
	}

	if result.WindowRect.Width > result.WindowRect.Height {
		_, err := s.Conn.RunSwayCommand(fmt.Sprintf("[con_id=%d] split h", event.Container.ID))
		return err
	} else {
		_, err := s.Conn.RunSwayCommand(fmt.Sprintf("[con_id=%d] split v", event.Container.ID))
		return err
	}

	return nil
}

// Manage makes layout is manageable via swaymgr
func (s *SpiralLayout) Manage() error {
	ws, err := s.Conn.GetFocusedWorkspace()
	if err != nil {
		return err
	}

	wc := WorkspaceConfig{
		Name:    ws.Name,
		Layout:  "spiral",
		Managed: true,
	}

	if err := s.store.put([]byte(ws.Name), wc); err != nil {
		return err
	}

	return nil
}
