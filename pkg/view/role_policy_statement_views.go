// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// RolePolicyStatementInventoryView RolePolicyStatement
type RolePolicyStatementInventoryView struct {
	Uuid       string              `json:"uuid,omitempty"`
	CreateDate time.Time           `json:"createDate,omitempty"`
	LastOpDate time.Time           `json:"lastOpDate,omitempty"`
	Statement  PolicyStatementView `json:"statement,omitempty"`
}
