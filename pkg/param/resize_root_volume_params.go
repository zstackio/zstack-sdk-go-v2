// Copyright (c) ZStack.io, Inc.

package param

// ResizeRootVolumeDetailParam ResizeRootVolume detail param
type ResizeRootVolumeDetailParam struct {
	Uuid string `json:"uuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	Size int64 `json:"size" validate:"required"`
}

// ResizeRootVolumeParam ResizeRootVolume request param
type ResizeRootVolumeParam struct {
	BaseParam
	Params ResizeRootVolumeDetailParam `json:"params"`
}
