// Copyright (c) ZStack.io, Inc.

package view

import "time"

// RoleInventoryView Role
type RoleInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"identity,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"state,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []RolePolicyStatementInventoryView `json:"statements,omitempty"`
	rest []RolePolicyRefInventoryView `json:"policies,omitempty"`
}

