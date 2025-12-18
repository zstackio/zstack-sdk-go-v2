// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2VirtualIDGroupInventoryView IAM2VirtualIDGroup
type IAM2VirtualIDGroupInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"projectUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []IAM2AttributeInventoryView `json:"attributes,omitempty"`
}

