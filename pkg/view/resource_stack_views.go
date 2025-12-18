// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ResourceStackInventoryView ResourceStack
type ResourceStackInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Version string `json:"version,omitempty"`
	Type string `json:"type,omitempty"`
	TemplateContent string `json:"templateContent,omitempty"`
	ParamContent string `json:"paramContent,omitempty"`
	Status string `json:"status,omitempty"`
	Reason string `json:"reason,omitempty"`
	Outputs string `json:"outputs,omitempty"`
	EnableRollback bool `json:"enableRollback,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

