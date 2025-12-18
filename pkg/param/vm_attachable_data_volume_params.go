// Copyright (c) ZStack.io, Inc.

package param

// GetVmAttachableDataVolumeDetailParam GetVmAttachableDataVolume详细参数
type GetVmAttachableDataVolumeDetailParam struct {
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// GetVmAttachableDataVolumeParam GetVmAttachableDataVolume请求参数
type GetVmAttachableDataVolumeParam struct {
	BaseParam
	Params GetVmAttachableDataVolumeDetailParam `json:"params"` // 详细参数
}

