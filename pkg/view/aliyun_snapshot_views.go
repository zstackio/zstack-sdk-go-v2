// Copyright (c) ZStack.io, Inc.

package view

import "time"

// AliyunSnapshotInventoryView AliyunSnapshot
type AliyunSnapshotInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"snapshotId,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"dataCenterUuid,omitempty"`
	rest string `json:"diskUuid,omitempty"`
	rest string `json:"status,omitempty"`
	rest string `json:"aliyunSnapshotUsage,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

