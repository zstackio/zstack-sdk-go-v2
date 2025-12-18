// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2VirtualIDGroupRoleRefInventoryView IAM2VirtualIDGroupRoleRef
type IAM2VirtualIDGroupRoleRefInventoryView struct {
	rest string `json:"groupUuid,omitempty"`
	rest string `json:"roleUuid,omitempty"`
	rest string `json:"targetAccountUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

