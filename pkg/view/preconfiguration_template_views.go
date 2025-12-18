// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// PreconfigurationTemplateInventoryView PreconfigurationTemplate
type PreconfigurationTemplateInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Distribution string `json:"distribution,omitempty"`
	Type string `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
	Md5sum string `json:"md5sum,omitempty"`
	IsPredefined bool `json:"isPredefined,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	CustomParams []string `json:"customParams,omitempty"`
}

