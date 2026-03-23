// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2ProjectRoleInventoryView IAM2ProjectRole
type IAM2ProjectRoleInventoryView struct {
	BaseInfoView
	BaseTimeView
	Iam2ProjectRoleType string `json:"iam2ProjectRoleType,omitempty"`
	Description string `json:"description,omitempty"`
	Identity string `json:"identity,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Statements []RolePolicyStatementInventoryView `json:"statements,omitempty"`
	Policies []RolePolicyRefInventoryView `json:"policies,omitempty"`
}

// CreateRoleEventView CreateRoleEvent
type CreateRoleEventView struct {
	Inventory RoleInventoryView `json:"inventory,omitempty"`
}

// QueryIAM2ProjectRoleView QueryIAM2ProjectRole
type QueryIAM2ProjectRoleView struct {
	Inventories []IAM2ProjectRoleInventoryView `json:"inventories,omitempty"`
}

