// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// AliyunPanguPartitionInventoryView AliyunPanguPartition
type AliyunPanguPartitionInventoryView struct {
	AccountUuid string `json:"accountUuid,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	IdentityZoneUuid string `json:"identityZoneUuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AppName string `json:"appName,omitempty"`
	PartitionName string `json:"partitionName,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
}

