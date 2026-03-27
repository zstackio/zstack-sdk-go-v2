// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// EcsVSwitchInventoryView EcsVSwitch
type EcsVSwitchInventoryView struct {
	BaseInfoView
	BaseTimeView
	VSwitchId string `json:"vSwitchId,omitempty"`
	Status string `json:"status,omitempty"`
	CidrBlock string `json:"cidrBlock,omitempty"`
	AvailableIpAddressCount int `json:"availableIpAddressCount,omitempty"`
	Description string `json:"description,omitempty"`
	EcsVpcUuid string `json:"ecsVpcUuid,omitempty"`
	IdentityZoneUuid string `json:"identityZoneUuid,omitempty"`
}

// SyncEcsVSwitchFromRemoteEventView SyncEcsVSwitchFromRemoteEvent
type SyncEcsVSwitchFromRemoteEventView struct {
	Inventories []EcsVSwitchInventoryView `json:"inventories,omitempty"`
}

// UpdateEcsVSwitchEventView UpdateEcsVSwitchEvent
type UpdateEcsVSwitchEventView struct {
	Inventory EcsVSwitchInventoryView `json:"inventory,omitempty"`
}

// CreateEcsVSwitchRemoteEventView CreateEcsVSwitchRemoteEvent
type CreateEcsVSwitchRemoteEventView struct {
	Inventory EcsVSwitchInventoryView `json:"inventory,omitempty"`
}

// QueryEcsVSwitchFromLocalView QueryEcsVSwitchFromLocal
type QueryEcsVSwitchFromLocalView struct {
	Inventories []EcsVSwitchInventoryView `json:"inventories,omitempty"`
}

