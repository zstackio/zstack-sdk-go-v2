// Copyright (c) ZStack.io, Inc.

package param

// DetachL3NetworkFromVmDetailParam DetachL3NetworkFromVm详细参数
type DetachL3NetworkFromVmDetailParam struct {
	rest string `json:"vmNicUuid" validate:"required"` // 必填
}

// DetachL3NetworkFromVmParam DetachL3NetworkFromVm请求参数
type DetachL3NetworkFromVmParam struct {
	BaseParam
	Params DetachL3NetworkFromVmDetailParam `json:"params"` // 详细参数
}

