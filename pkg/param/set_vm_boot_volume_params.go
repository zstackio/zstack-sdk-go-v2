// Copyright (c) ZStack.io, Inc.

package param

// SetVmBootVolumeDetailParam SetVmBootVolume详细参数
type SetVmBootVolumeDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"volumeUuid" validate:"required"` // 必填
}

// SetVmBootVolumeParam SetVmBootVolume请求参数
type SetVmBootVolumeParam struct {
	BaseParam
	Params SetVmBootVolumeDetailParam `json:"params"` // 详细参数
}

