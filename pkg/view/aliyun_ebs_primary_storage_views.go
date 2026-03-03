// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunEbsPrimaryStorageInventoryView AliyunEbsPrimaryStorage
type AliyunEbsPrimaryStorageInventoryView struct {
	BaseInfoView
	BaseTimeView
	PanguAppName string `json:"panguAppName,omitempty"`
	PanguPartitionName string `json:"panguPartitionName,omitempty"`
	IdentityZoneUuid string `json:"identityZoneUuid,omitempty"`
	DefaultIoType string `json:"defaultIoType,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	TotalPhysicalCapacity int64 `json:"totalPhysicalCapacity,omitempty"`
	AvailablePhysicalCapacity int64 `json:"availablePhysicalCapacity,omitempty"`
	SystemUsedCapacity int64 `json:"systemUsedCapacity,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	MountPath string `json:"mountPath,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
}

