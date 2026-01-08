// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EcsSecurityGroupInventoryView EcsSecurityGroup
type EcsSecurityGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	EcsVpcUuid      string `json:"ecsVpcUuid,omitempty"`
	SecurityGroupId string `json:"securityGroupId,omitempty"`
}

// SyncEcsSecurityGroupFromRemoteEventView SyncEcsSecurityGroupFromRemoteEvent
type SyncEcsSecurityGroupFromRemoteEventView struct {
	Inventories []EcsSecurityGroupInventoryView `json:"inventories,omitempty"`
}

// UpdateEcsSecurityGroupEventView UpdateEcsSecurityGroupEvent
type UpdateEcsSecurityGroupEventView struct {
	Inventory EcsSecurityGroupInventoryView `json:"inventory,omitempty"`
}

// QueryEcsSecurityGroupFromLocalView QueryEcsSecurityGroupFromLocal
type QueryEcsSecurityGroupFromLocalView struct {
	Inventories []EcsSecurityGroupInventoryView `json:"inventories,omitempty"`
}

// CreateEcsSecurityGroupRemoteEventView CreateEcsSecurityGroupRemoteEvent
type CreateEcsSecurityGroupRemoteEventView struct {
	Inventory EcsSecurityGroupInventoryView `json:"inventory,omitempty"`
}
