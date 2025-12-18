// Copyright (c) ZStack.io, Inc.

package param

// DetachDataVolumeFromHostDetailParam DetachDataVolumeFromHost详细参数
type DetachDataVolumeFromHostDetailParam struct {
	rest string `json:"volumeUuid" validate:"required"` // 必填
	rest string `json:"hostUuid,omitempty"`
}

// DetachDataVolumeFromHostParam DetachDataVolumeFromHost请求参数
type DetachDataVolumeFromHostParam struct {
	BaseParam
	Params DetachDataVolumeFromHostDetailParam `json:"params"` // 详细参数
}

