// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LogServerInventoryView LogServer
type LogServerInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"category,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"level,omitempty"`
	rest string `json:"configuration,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

