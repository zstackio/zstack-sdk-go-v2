// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AliyunDiskInventoryView AliyunDisk
type AliyunDiskInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"diskId,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"identityZoneUuid,omitempty"`
	rest string `json:"ecsInstanceUuid,omitempty"`
	rest string `json:"diskCategory,omitempty"`
	rest string `json:"diskType,omitempty"`
	rest string `json:"diskChargeType,omitempty"`
	rest string `json:"status,omitempty"`
	rest int `json:"sizeWithGB,omitempty"`
	rest string `json:"deviceInfo,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

