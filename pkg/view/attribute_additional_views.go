// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AttributeView Attribute
type AttributeView struct {
	Uuid  string `json:"uuid,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
