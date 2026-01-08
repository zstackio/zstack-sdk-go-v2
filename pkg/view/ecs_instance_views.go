// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EcsInstanceInventoryView EcsInstance
type EcsInstanceInventoryView struct {
	BaseInfoView
	BaseTimeView
	LocalVmInstanceUuid   string    `json:"localVmInstanceUuid,omitempty"`
	EcsInstanceId         string    `json:"ecsInstanceId,omitempty"`
	EcsStatus             string    `json:"ecsStatus,omitempty"`
	CpuCores              int64     `json:"cpuCores,omitempty"`
	MemorySize            int64     `json:"memorySize,omitempty"`
	EcsInstanceType       string    `json:"ecsInstanceType,omitempty"`
	EcsBandWidth          int64     `json:"ecsBandWidth,omitempty"`
	EcsRootVolumeId       string    `json:"ecsRootVolumeId,omitempty"`
	EcsRootVolumeCategory string    `json:"ecsRootVolumeCategory,omitempty"`
	EcsRootVolumeSize     int64     `json:"ecsRootVolumeSize,omitempty"`
	PrivateIpAddress      string    `json:"privateIpAddress,omitempty"`
	PublicIpAddress       string    `json:"publicIpAddress,omitempty"`
	EcsVSwitchUuid        string    `json:"ecsVSwitchUuid,omitempty"`
	EcsImageUuid          string    `json:"ecsImageUuid,omitempty"`
	EcsSecurityGroupUuid  string    `json:"ecsSecurityGroupUuid,omitempty"`
	IdentityZoneUuid      string    `json:"identityZoneUuid,omitempty"`
	ChargeType            string    `json:"chargeType,omitempty"`
	ExpireDate            time.Time `json:"expireDate,omitempty"`
}

// StartEcsInstanceEventView StartEcsInstanceEvent
type StartEcsInstanceEventView struct {
	Success bool `json:"success,omitempty"`
}

// DeleteEcsInstanceEventView DeleteEcsInstanceEvent
type DeleteEcsInstanceEventView struct {
	Success bool `json:"success,omitempty"`
}

// StopEcsInstanceEventView StopEcsInstanceEvent
type StopEcsInstanceEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryEcsInstanceFromLocalView QueryEcsInstanceFromLocal
type QueryEcsInstanceFromLocalView struct {
	Inventories []EcsInstanceInventoryView `json:"inventories,omitempty"`
}

// RebootEcsInstanceEventView RebootEcsInstanceEvent
type RebootEcsInstanceEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateEcsInstanceEventView UpdateEcsInstanceEvent
type UpdateEcsInstanceEventView struct {
	Inventory EcsInstanceInventoryView `json:"inventory,omitempty"`
}

// SyncEcsInstanceFromRemoteEventView SyncEcsInstanceFromRemoteEvent
type SyncEcsInstanceFromRemoteEventView struct {
	Inventories []EcsInstanceInventoryView `json:"inventories,omitempty"`
}

// CreateEcsInstanceFromEcsImageEventView CreateEcsInstanceFromEcsImageEvent
type CreateEcsInstanceFromEcsImageEventView struct {
	Inventory EcsInstanceInventoryView `json:"inventory,omitempty"`
}
