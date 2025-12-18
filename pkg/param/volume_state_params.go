// Copyright (c) ZStack.io, Inc.

package param

// ChangeVolumeStateDetailParam ChangeVolumeState详细参数
type ChangeVolumeStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeVolumeStateParam ChangeVolumeState请求参数
type ChangeVolumeStateParam struct {
	BaseParam
	Params ChangeVolumeStateDetailParam `json:"params"` // 详细参数
}

