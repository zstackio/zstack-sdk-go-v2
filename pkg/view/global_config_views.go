// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// GlobalConfigInventoryView GlobalConfig
type GlobalConfigInventoryView struct {
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	DefaultValue string `json:"defaultValue,omitempty"`
	Value string `json:"value,omitempty"`
}

