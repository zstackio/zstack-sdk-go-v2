// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateBackupStorageForCreatingImageDetailParam GetCandidateBackupStorageForCreatingImage详细参数
type GetCandidateBackupStorageForCreatingImageDetailParam struct {
	rest string `json:"volumeUuid,omitempty"`
	rest string `json:"volumeSnapshotUuid,omitempty"`
}

// GetCandidateBackupStorageForCreatingImageParam GetCandidateBackupStorageForCreatingImage请求参数
type GetCandidateBackupStorageForCreatingImageParam struct {
	BaseParam
	Params GetCandidateBackupStorageForCreatingImageDetailParam `json:"params"` // 详细参数
}

