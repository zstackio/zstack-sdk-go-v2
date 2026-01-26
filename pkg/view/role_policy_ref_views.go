// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// RolePolicyRefInventoryView RolePolicyRef
type RolePolicyRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	RoleUuid string `json:"roleUuid,omitempty"`
	PolicyUuid string `json:"policyUuid,omitempty"`
}

