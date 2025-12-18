// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SecurityGroupL3NetworkRefInventoryView SecurityGroupL3NetworkRef
type SecurityGroupL3NetworkRefInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"l3NetworkUuid,omitempty"`
	rest string `json:"securityGroupUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

