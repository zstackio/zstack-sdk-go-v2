// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ZdfsStorageInventoryView ZdfsStorage
type ZdfsStorageInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"endPoint,omitempty"`
	rest string `json:"accessKey,omitempty"`
	rest string `json:"secretKey,omitempty"`
	rest string `json:"type,omitempty"`
	rest int64 `json:"usedCapacity,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

