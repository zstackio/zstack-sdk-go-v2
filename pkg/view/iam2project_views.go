// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2ProjectInventoryView IAM2Project
type IAM2ProjectInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []IAM2AttributeInventoryView `json:"attributes,omitempty"`
	rest string `json:"linkedAccountUuid,omitempty"`
}

