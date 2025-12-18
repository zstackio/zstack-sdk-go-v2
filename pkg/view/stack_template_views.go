// Copyright (c) ZStack.io, Inc.

package view

import "time"

// StackTemplateInventoryView StackTemplate
type StackTemplateInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"version,omitempty"`
	rest bool `json:"state,omitempty"`
	rest string `json:"content,omitempty"`
	rest string `json:"md5sum,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

