// Copyright (c) ZStack.io, Inc.

package param

// ResizeDataVolumeDetailParam ResizeDataVolume detail param
type ResizeDataVolumeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Size int64 `json:"size" validate:"required"`
}

// ResizeDataVolumeParam ResizeDataVolume request param
type ResizeDataVolumeParam struct {
	BaseParam
	Params ResizeDataVolumeDetailParam `json:"params"`
}
