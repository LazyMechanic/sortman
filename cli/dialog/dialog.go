package dialog

import "github.com/LazyMechanic/sortman/types"

const (
	AddRequest   string       = "Add request"
	ShowRequests string       = "Show requests"
	Execute      string       = "Execute"
	Cancel       string       = "Cancel"
	CopyAction   types.Action = "copy"
	MoveAction   types.Action = "move"
)

var (
	WhatToDo = []string{
		AddRequest,
		ShowRequests,
		Execute,
		Cancel,
	}
)
