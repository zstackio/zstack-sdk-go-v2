// Copyright (c) ZStack.io, Inc.

package param

// AttachDataVolumeToVmDetailParam AttachDataVolumeToVm详细参数
type AttachDataVolumeToVmDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest string `json:"volumeUuid" validate:"required"` // 必填
}

// AttachDataVolumeToVmParam AttachDataVolumeToVm请求参数
type AttachDataVolumeToVmParam struct {
	BaseParam
	Params AttachDataVolumeToVmDetailParam `json:"params"` // 详细参数
}

