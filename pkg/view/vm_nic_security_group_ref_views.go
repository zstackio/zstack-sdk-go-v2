// Copyright (c) ZStack.io, Inc.

package view

import "time"

// VmNicSecurityGroupRefInventoryView VmNicSecurityGroupRef
type VmNicSecurityGroupRefInventoryView struct {
	rest int `json:"priority,omitempty"`
	rest string `json:"vmNicUuid,omitempty"`
	rest string `json:"securityGroupUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

