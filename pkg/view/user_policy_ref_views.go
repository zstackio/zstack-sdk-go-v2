// Copyright (c) ZStack.io, Inc.

package view

import "time"

// UserPolicyRefInventoryView UserPolicyRef
type UserPolicyRefInventoryView struct {
	rest string `json:"userUuid,omitempty"`
	rest string `json:"policyUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

