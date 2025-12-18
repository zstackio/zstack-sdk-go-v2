// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VxlanHostMappingInventoryView VxlanHostMapping
type VxlanHostMappingInventoryView struct {
	VxlanUuid string `json:"vxlanUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	VlanId int `json:"vlanId,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

