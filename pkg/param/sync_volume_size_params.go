// Copyright (c) ZStack.io, Inc.

package param

// SyncVolumeSizeDetailParam SyncVolumeSize detail param
type SyncVolumeSizeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// SyncVolumeSizeParam SyncVolumeSize request param
type SyncVolumeSizeParam struct {
	BaseParam
	Params SyncVolumeSizeDetailParam `json:"params"`
}
