// Copyright (c) ZStack.io, Inc.

package param

// GetDataVolumeAttachableVmDetailParam GetDataVolumeAttachableVm详细参数
type GetDataVolumeAttachableVmDetailParam struct {
	rest string `json:"volumeUuid" validate:"required"` // 必填
}

// GetDataVolumeAttachableVmParam GetDataVolumeAttachableVm请求参数
type GetDataVolumeAttachableVmParam struct {
	BaseParam
	Params GetDataVolumeAttachableVmDetailParam `json:"params"` // 详细参数
}

