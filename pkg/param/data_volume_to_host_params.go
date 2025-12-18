// Copyright (c) ZStack.io, Inc.

package param

// AttachDataVolumeToHostDetailParam AttachDataVolumeToHost详细参数
type AttachDataVolumeToHostDetailParam struct {
	rest string `json:"volumeUuid" validate:"required"` // 必填
	rest string `json:"hostUuid" validate:"required"` // 必填
	rest string `json:"mountPath" validate:"required"` // 必填
}

// AttachDataVolumeToHostParam AttachDataVolumeToHost请求参数
type AttachDataVolumeToHostParam struct {
	BaseParam
	Params AttachDataVolumeToHostDetailParam `json:"params"` // 详细参数
}

