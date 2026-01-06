// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// UserPolicyRefInventoryView UserPolicyRef
type UserPolicyRefInventoryView struct {
	UserUuid string `json:"userUuid,omitempty"`
	PolicyUuid string `json:"policyUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

