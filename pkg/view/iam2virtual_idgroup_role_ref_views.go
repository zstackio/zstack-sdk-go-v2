// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IAM2VirtualIDGroupRoleRefInventoryView IAM2VirtualIDGroupRoleRef
type IAM2VirtualIDGroupRoleRefInventoryView struct {
	GroupUuid string `json:"groupUuid,omitempty"`
	RoleUuid string `json:"roleUuid,omitempty"`
	TargetAccountUuid string `json:"targetAccountUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

