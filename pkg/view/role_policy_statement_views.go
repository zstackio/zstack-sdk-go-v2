// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// RolePolicyStatementInventoryView RolePolicyStatement
type RolePolicyStatementInventoryView struct {
	BaseInfoView
	BaseTimeView
	Statement PolicyStatementView `json:"statement,omitempty"`
}

