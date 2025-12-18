// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ResourceStackInventoryView ResourceStack
type ResourceStackInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"version,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"templateContent,omitempty"`
	rest string `json:"paramContent,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"reason,omitempty"`
	rest string `json:"outputs,omitempty"`
	rest bool `json:"enableRollback,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

