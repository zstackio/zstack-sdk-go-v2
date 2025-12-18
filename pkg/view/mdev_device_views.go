// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// MdevDeviceInventoryView MdevDevice
type MdevDeviceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	MttyUuid string `json:"mttyUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	MdevSpecUuid string `json:"mdevSpecUuid,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Chooser string `json:"chooser,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Vendor string `json:"vendor,omitempty"`
}

