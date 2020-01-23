package main

import (
	"fmt"

	"github.com/Difrex/gosway/ipc"
)

type LeftLayout struct {
	Conn  *ipc.SwayConnection
	store *store
}

// NewLeftLayout initializes the LeftLayout interface
func NewLeftLayout(conn *ipc.SwayConnection, store *store) *LeftLayout {
	layout := &LeftLayout{}
	layout.Conn = conn
	layout.store = store
	return layout
}

func (l *LeftLayout) OnFocus(event *ipc.Event) error {
	return nil
}

// PlaceWindow ...
func (l *LeftLayout) OnNew(event *ipc.Event) error {
	nodes, err := l.Conn.GetFocusedWorkspaceWindows()
	if err != nil {
		return err
	}

	// Initial places
	if len(nodes) <= 1 {
		if _, err := l.Conn.RunSwayCommand(fmt.Sprintf("[con_id=%d] split h", event.Container.ID)); err != nil {
			return err
		}
	} else if len(nodes) == 2 {
		if _, err := l.Conn.RunSwayCommand(fmt.Sprintf("[con_id=%d] split v", event.Container.ID)); err != nil {
			return err
		}
	} else {
		// n, err := l.Conn.GetFocusedWorkspaceWindows()
		// if err != nil {
		// 	return err
		// }
		// ch := make(chan ipc.Node)
		// go ipc.FindFocusedNodes(n, ch)
		// node := <-ch

		// fmt.Println("Finded id: ", node.ID)
		// fmt.Println("Event id: ", event.Container.ID)
		// fmt.Println("Largest: ", ipc.GetLargestWindowID(nodes))
		// fmt.Printf("%+v", event.Container)
		// if node.ID != ipc.GetLargestWindowID(nodes) {
		// 	o, err := l.Conn.RunSwayCommand(fmt.Sprintf("[con_id=%d] split v", event.Container.ID))
		// 	if err != nil {
		// 		return err
		// 	}
		// 	fmt.Println(string(o))
		// 	o2, err := l.Conn.RunSwayCommand(fmt.Sprintf("[con_id=%d] move right", event.Container.ID))
		// 	if err != nil {
		// 		return err
		// 	}
		// 	fmt.Println(string(o2))
		// } else {
		// 	wfmt.Println("smth went wrong")
		// }
	}

	return nil
}

// Manage makes layout is manageable via swaymgr
func (l *LeftLayout) Manage() error {
	ws, err := l.Conn.GetFocusedWorkspace()
	if err != nil {
		return err
	}

	wc := WorkspaceConfig{
		Name:    ws.Name,
		Layout:  "left",
		Managed: true,
	}

	return l.store.put([]byte(ws.Name), wc)
}
