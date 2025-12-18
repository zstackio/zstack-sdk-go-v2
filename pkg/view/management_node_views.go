// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ManagementNodeInventoryView ManagementNode
type ManagementNodeInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"hostName,omitempty"`
	rest time.Time `json:"joinDate,omitempty"`
	rest time.Time `json:"heartBeat,omitempty"`
}

