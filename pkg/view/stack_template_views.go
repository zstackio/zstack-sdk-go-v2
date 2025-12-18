// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// StackTemplateInventoryView StackTemplate
type StackTemplateInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	Version string `json:"version,omitempty"`
	State bool `json:"state,omitempty"`
	Content string `json:"content,omitempty"`
	Md5sum string `json:"md5sum,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

