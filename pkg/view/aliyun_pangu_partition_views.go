// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AliyunPanguPartitionInventoryView AliyunPanguPartition
type AliyunPanguPartitionInventoryView struct {
	rest string `json:"accountUuid,omitempty"`
	rest string `json:"uuid,omitempty"`
	rest string `json:"identityZoneUuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"appName,omitempty"`
	rest string `json:"partitionName,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

