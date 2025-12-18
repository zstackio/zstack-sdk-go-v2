// Copyright (c) ZStack.io, Inc.

package view

import "time"

// UserGroupPolicyRefInventoryView UserGroupPolicyRef
type UserGroupPolicyRefInventoryView struct {
	rest string `json:"groupUuid,omitempty"`
	rest string `json:"policyUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

