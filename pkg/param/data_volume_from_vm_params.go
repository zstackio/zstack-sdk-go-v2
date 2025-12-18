// Copyright (c) ZStack.io, Inc.

package param

// DetachDataVolumeFromVmDetailParam DetachDataVolumeFromVm详细参数
type DetachDataVolumeFromVmDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"vmUuid,omitempty"`
}

// DetachDataVolumeFromVmParam DetachDataVolumeFromVm请求参数
type DetachDataVolumeFromVmParam struct {
	BaseParam
	Params DetachDataVolumeFromVmDetailParam `json:"params"` // 详细参数
}

