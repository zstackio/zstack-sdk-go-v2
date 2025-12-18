// Copyright (c) ZStack.io, Inc.

package view

// BatchSyncVolumeSizeView BatchSyncVolumeSize
type BatchSyncVolumeSizeView struct {
	SuccessCount int `json:"successCount,omitempty"`
	FailCount int `json:"failCount,omitempty"`
	Success bool `json:"success,omitempty"`
}

