// Copyright (c) ZStack.io, Inc.

package view

import "time"

// RolePolicyRefInventoryView RolePolicyRef
type RolePolicyRefInventoryView struct {
	rest string `json:"roleUuid,omitempty"`
	rest string `json:"policyUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

