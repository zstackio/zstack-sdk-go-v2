// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// EcsInstanceInventoryView EcsInstance
type EcsInstanceInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	LocalVmInstanceUuid string `json:"localVmInstanceUuid,omitempty"`
	EcsInstanceId string `json:"ecsInstanceId,omitempty"`
	Name string `json:"name,omitempty"`
	EcsStatus string `json:"ecsStatus,omitempty"`
	CpuCores int64 `json:"cpuCores,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	EcsInstanceType string `json:"ecsInstanceType,omitempty"`
	EcsBandWidth int64 `json:"ecsBandWidth,omitempty"`
	EcsRootVolumeId string `json:"ecsRootVolumeId,omitempty"`
	EcsRootVolumeCategory string `json:"ecsRootVolumeCategory,omitempty"`
	EcsRootVolumeSize int64 `json:"ecsRootVolumeSize,omitempty"`
	PrivateIpAddress string `json:"privateIpAddress,omitempty"`
	PublicIpAddress string `json:"publicIpAddress,omitempty"`
	EcsVSwitchUuid string `json:"ecsVSwitchUuid,omitempty"`
	EcsImageUuid string `json:"ecsImageUuid,omitempty"`
	EcsSecurityGroupUuid string `json:"ecsSecurityGroupUuid,omitempty"`
	IdentityZoneUuid string `json:"identityZoneUuid,omitempty"`
	ChargeType string `json:"chargeType,omitempty"`
	ExpireDate time.Time `json:"expireDate,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	Description string `json:"description,omitempty"`
}

