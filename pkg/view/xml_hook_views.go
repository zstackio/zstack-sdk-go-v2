// Copyright (c) ZStack.io, Inc.

package view

import "time"

// XmlHookInventoryView XmlHook
type XmlHookInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"hookScript,omitempty"`
	rest string `json:"libvirtVersion,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

