// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IAM2VirtualIDRoleRefInventoryView IAM2VirtualIDRoleRef
type IAM2VirtualIDRoleRefInventoryView struct {
	rest string `json:"virtualIDUuid,omitempty"`
	rest string `json:"roleUuid,omitempty"`
	rest string `json:"targetAccountUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

