// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ZdfsStorageInventoryView ZdfsStorage
type ZdfsStorageInventoryView struct {
	BaseInfoView
	BaseTimeView
	EndPoint string `json:"endPoint,omitempty"`
	AccessKey string `json:"accessKey,omitempty"`
	SecretKey string `json:"secretKey,omitempty"`
	Type string `json:"type,omitempty"`
	UsedCapacity int64 `json:"usedCapacity,omitempty"`
}

