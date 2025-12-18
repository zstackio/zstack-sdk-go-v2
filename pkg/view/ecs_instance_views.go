// Copyright (c) ZStack.io, Inc.

package view

import "time"

// EcsInstanceInventoryView EcsInstance
type EcsInstanceInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"localVmInstanceUuid,omitempty"`
	rest string `json:"ecsInstanceId,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"ecsStatus,omitempty"`
	rest int64 `json:"cpuCores,omitempty"`
	rest int64 `json:"memorySize,omitempty"`
	rest string `json:"ecsInstanceType,omitempty"`
	rest int64 `json:"ecsBandWidth,omitempty"`
	rest string `json:"ecsRootVolumeId,omitempty"`
	rest string `json:"ecsRootVolumeCategory,omitempty"`
	rest int64 `json:"ecsRootVolumeSize,omitempty"`
	rest string `json:"privateIpAddress,omitempty"`
	rest string `json:"publicIpAddress,omitempty"`
	rest string `json:"ecsVSwitchUuid,omitempty"`
	rest string `json:"ecsImageUuid,omitempty"`
	rest string `json:"ecsSecurityGroupUuid,omitempty"`
	rest string `json:"identityZoneUuid,omitempty"`
	rest string `json:"chargeType,omitempty"`
	rest time.Time `json:"expireDate,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest string `json:"description,omitempty"`
}

