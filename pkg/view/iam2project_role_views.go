// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2ProjectRoleInventoryView IAM2ProjectRole
type IAM2ProjectRoleInventoryView struct {
	Iam2ProjectRoleType string `json:"iam2ProjectRoleType,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Identity string `json:"identity,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Statements []RolePolicyStatementInventoryView `json:"statements,omitempty"`
	Policies []RolePolicyRefInventoryView `json:"policies,omitempty"`
}

