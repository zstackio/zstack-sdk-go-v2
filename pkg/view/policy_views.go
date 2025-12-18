// Copyright (c) ZStack.io, Inc.

package view

import "time"

// PolicyView Policy
type PolicyView struct {
	rest interface{} `json:"metaClass,omitempty"`
	rest []interface{} `json:"statements,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"accountUuid,omitempty"`
}

// PolicyInventoryView Policy
type PolicyInventoryView struct {
	rest []interface{} `json:"statements,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"accountUuid,omitempty"`
}

