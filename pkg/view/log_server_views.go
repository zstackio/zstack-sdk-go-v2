// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LogServerInventoryView LogServer
type LogServerInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Category string `json:"category,omitempty"`
	Type string `json:"type,omitempty"`
	Level string `json:"level,omitempty"`
	Configuration string `json:"configuration,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

