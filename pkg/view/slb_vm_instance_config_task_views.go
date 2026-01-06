// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SlbVmInstanceConfigTaskInventoryView SlbVmInstanceConfigTask
type SlbVmInstanceConfigTaskInventoryView struct {
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	ConfigVersion int64 `json:"configVersion,omitempty"`
	TaskName string `json:"taskName,omitempty"`
	TaskData string `json:"taskData,omitempty"`
	RetryNumber int64 `json:"retryNumber,omitempty"`
	LastFailedReason string `json:"lastFailedReason,omitempty"`
	Status string `json:"status,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

