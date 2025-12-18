// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2ProjectTemplateInventoryView IAM2ProjectTemplate
type IAM2ProjectTemplateInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest interface{} `json:"template,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

