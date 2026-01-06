// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// UserGroupPolicyRefInventoryView UserGroupPolicyRef
type UserGroupPolicyRefInventoryView struct {
	GroupUuid string `json:"groupUuid,omitempty"`
	PolicyUuid string `json:"policyUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

