// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// RolePolicyRefInventoryView RolePolicyRef
type RolePolicyRefInventoryView struct {
	RoleUuid   string    `json:"roleUuid,omitempty"`
	PolicyUuid string    `json:"policyUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}
