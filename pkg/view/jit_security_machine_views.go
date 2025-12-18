// Copyright (c) ZStack.io, Inc.

package view

import "time"

// JitSecurityMachineInventoryView JitSecurityMachine
type JitSecurityMachineInventoryView struct {
	rest int `json:"port,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"secretResourcePoolUuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"managementIp,omitempty"`
	rest string `json:"type,omitempty"`
	rest string `json:"model,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

