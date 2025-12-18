// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VmNicSecurityGroupRefInventoryView VmNicSecurityGroupRef
type VmNicSecurityGroupRefInventoryView struct {
	Priority int `json:"priority,omitempty"`
	VmNicUuid string `json:"vmNicUuid,omitempty"`
	SecurityGroupUuid string `json:"securityGroupUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

