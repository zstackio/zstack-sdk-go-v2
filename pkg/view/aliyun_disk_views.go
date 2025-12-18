// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunDiskInventoryView AliyunDisk
type AliyunDiskInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	DiskId string `json:"diskId,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IdentityZoneUuid string `json:"identityZoneUuid,omitempty"`
	EcsInstanceUuid string `json:"ecsInstanceUuid,omitempty"`
	DiskCategory string `json:"diskCategory,omitempty"`
	DiskType string `json:"diskType,omitempty"`
	DiskChargeType string `json:"diskChargeType,omitempty"`
	Status string `json:"status,omitempty"`
	SizeWithGB int `json:"sizeWithGB,omitempty"`
	DeviceInfo string `json:"deviceInfo,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

