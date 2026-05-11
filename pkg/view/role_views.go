// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// RoleInventoryView Role
type RoleInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Identity string `json:"identity,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Statements []RolePolicyStatementInventoryView `json:"statements,omitempty"`
	Policies []RolePolicyRefInventoryView `json:"policies,omitempty"`
}

// ChangeRoleStateEventView ChangeRoleStateEvent
type ChangeRoleStateEventView struct {
	Inventory RoleInventoryView `json:"inventory,omitempty"`
}

// CreateRoleEventView CreateRoleEvent
type CreateRoleEventView struct {
	Inventory RoleInventoryView `json:"inventory,omitempty"`
}

// QueryRoleView QueryRole
type QueryRoleView struct {
	Inventories []RoleInventoryView `json:"inventories,omitempty"`
}

// DeleteRoleEventView DeleteRoleEvent
type DeleteRoleEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateRoleEventView UpdateRoleEvent
type UpdateRoleEventView struct {
	Inventory RoleInventoryView `json:"inventory,omitempty"`
}

