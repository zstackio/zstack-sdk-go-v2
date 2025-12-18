// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SlbVmInstanceConfigTaskInventoryView SlbVmInstanceConfigTask
type SlbVmInstanceConfigTaskInventoryView struct {
	rest string `json:"vmInstanceUuid,omitempty"`
	rest int64 `json:"configVersion,omitempty"`
	rest string `json:"taskName,omitempty"`
	rest string `json:"taskData,omitempty"`
	rest int64 `json:"retryNumber,omitempty"`
	rest string `json:"lastFailedReason,omitempty"`
	rest string `json:"status,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

