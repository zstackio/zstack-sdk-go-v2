// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// UserPolicyRefInventoryView UserPolicyRef
type UserPolicyRefInventoryView struct {
	UserUuid string `json:"userUuid,omitempty"`
	PolicyUuid string `json:"policyUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

