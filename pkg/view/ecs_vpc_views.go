// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EcsVpcInventoryView EcsVpc
type EcsVpcInventoryView struct {
	BaseInfoView
	BaseTimeView
	EcsVpcId *string `json:"ecsVpcId,omitempty"`
	DataCenterUuid *string `json:"dataCenterUuid,omitempty"`
	Status *string `json:"status,omitempty"`
	Deleted *string `json:"deleted,omitempty"`
	CidrBlock *string `json:"cidrBlock,omitempty"`
	VRouterId *string `json:"vRouterId,omitempty"`
	Description *string `json:"description,omitempty"`
}

// QueryEcsVpcFromLocalView QueryEcsVpcFromLocal
type QueryEcsVpcFromLocalView struct {
	Inventories []EcsVpcInventoryView `json:"inventories,omitempty"`
}

// CreateEcsVpcRemoteEventView CreateEcsVpcRemoteEvent
type CreateEcsVpcRemoteEventView struct {
	Inventory EcsVpcInventoryView `json:"inventory,omitempty"`
}

// UpdateEcsVpcEventView UpdateEcsVpcEvent
type UpdateEcsVpcEventView struct {
	Inventory EcsVpcInventoryView `json:"inventory,omitempty"`
}

// SyncEcsVpcFromRemoteEventView SyncEcsVpcFromRemoteEvent
type SyncEcsVpcFromRemoteEventView struct {
	Inventories []EcsVpcInventoryView `json:"inventories,omitempty"`
}

