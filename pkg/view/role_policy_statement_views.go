// Copyright (c) ZStack.io, Inc.

package view

import "time"

// RolePolicyStatementInventoryView RolePolicyStatement
type RolePolicyStatementInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest interface{} `json:"statement,omitempty"`
}

