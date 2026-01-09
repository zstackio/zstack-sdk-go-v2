// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// HardwareL2VxlanNetworkInventoryView HardwareL2VxlanNetwork
type HardwareL2VxlanNetworkInventoryView struct {
	Vlan *int `json:"vlan,omitempty"`
	Vni *int `json:"vni,omitempty"`
	PoolUuid *string `json:"poolUuid,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ZoneUuid *string `json:"zoneUuid,omitempty"`
	PhysicalInterface *string `json:"physicalInterface,omitempty"`
	Type *string `json:"type,omitempty"`
	VSwitchType *string `json:"vSwitchType,omitempty"`
	VirtualNetworkId *int `json:"virtualNetworkId,omitempty"`
	Isolated *bool `json:"isolated,omitempty"`
	Pvlan *string `json:"pvlan,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
}

