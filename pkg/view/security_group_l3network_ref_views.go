// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SecurityGroupL3NetworkRefInventoryView SecurityGroupL3NetworkRef
type SecurityGroupL3NetworkRefInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	SecurityGroupUuid string `json:"securityGroupUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

