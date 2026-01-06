// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// RolePolicyStatementInventoryView RolePolicyStatement
type RolePolicyStatementInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
	Statement PolicyStatementView `json:"statement,omitempty"`
}

