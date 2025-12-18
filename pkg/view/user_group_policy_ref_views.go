// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// UserGroupPolicyRefInventoryView UserGroupPolicyRef
type UserGroupPolicyRefInventoryView struct {
	GroupUuid string `json:"groupUuid,omitempty"`
	PolicyUuid string `json:"policyUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

