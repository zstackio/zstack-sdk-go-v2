// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// VtepInventoryView Vtep
type VtepInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	VtepIp string `json:"vtepIp,omitempty"`
	Port int `json:"port,omitempty"`
	Type string `json:"type,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	PoolUuid string `json:"poolUuid,omitempty"`
}

