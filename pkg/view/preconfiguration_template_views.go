// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PreconfigurationTemplateInventoryView PreconfigurationTemplate
type PreconfigurationTemplateInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"distribution,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"content,omitempty"`
	rest string `json:"md5sum,omitempty"`
	rest bool `json:"isPredefined,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []string `json:"customParams,omitempty"`
}

