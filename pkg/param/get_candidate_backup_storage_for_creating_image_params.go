// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateBackupStorageForCreatingImageDetailParam GetCandidateBackupStorageForCreatingImage detail param
type GetCandidateBackupStorageForCreatingImageDetailParam struct {
	VolumeUuid string `json:"volumeUuid,omitempty"`
	VolumeSnapshotUuid string `json:"volumeSnapshotUuid,omitempty"`
}

// GetCandidateBackupStorageForCreatingImageParam GetCandidateBackupStorageForCreatingImage request param
type GetCandidateBackupStorageForCreatingImageParam struct {
	BaseParam
	Params GetCandidateBackupStorageForCreatingImageDetailParam `json:"params"`
}
