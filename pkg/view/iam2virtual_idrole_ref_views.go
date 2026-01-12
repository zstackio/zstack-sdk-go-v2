// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2VirtualIDRoleRefInventoryView IAM2VirtualIDRoleRef
type IAM2VirtualIDRoleRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	VirtualIDUuid *string `json:"virtualIDUuid,omitempty"`
	RoleUuid *string `json:"roleUuid,omitempty"`
	TargetAccountUuid *string `json:"targetAccountUuid,omitempty"`
}

