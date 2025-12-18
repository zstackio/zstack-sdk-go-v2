// Copyright (c) ZStack.io, Inc.

package param

// SetVmSoundTypeDetailParam SetVmSoundType详细参数
type SetVmSoundTypeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"soundType" validate:"required"` // 必填
}

// SetVmSoundTypeParam SetVmSoundType请求参数
type SetVmSoundTypeParam struct {
	BaseParam
	Params SetVmSoundTypeDetailParam `json:"params"` // 详细参数
}

