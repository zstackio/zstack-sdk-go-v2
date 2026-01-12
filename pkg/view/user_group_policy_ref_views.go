// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// UserGroupPolicyRefInventoryView UserGroupPolicyRef
type UserGroupPolicyRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	GroupUuid *string `json:"groupUuid,omitempty"`
	PolicyUuid *string `json:"policyUuid,omitempty"`
}

