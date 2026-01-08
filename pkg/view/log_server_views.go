// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LogServerInventoryView LogServer
type LogServerInventoryView struct {
	BaseInfoView
	BaseTimeView
	Category      string `json:"category,omitempty"`
	Type          string `json:"type,omitempty"`
	Level         string `json:"level,omitempty"`
	Configuration string `json:"configuration,omitempty"`
}

// UpdateLogServerEventView UpdateLogServerEvent
type UpdateLogServerEventView struct {
	Inventory LogServerInventoryView `json:"inventory,omitempty"`
}

// DeleteLogServerEventView DeleteLogServerEvent
type DeleteLogServerEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryLogServerView QueryLogServer
type QueryLogServerView struct {
	Inventories []LogServerInventoryView `json:"inventories,omitempty"`
}

// AddLogServerEventView AddLogServerEvent
type AddLogServerEventView struct {
	Inventory LogServerInventoryView `json:"inventory,omitempty"`
}
