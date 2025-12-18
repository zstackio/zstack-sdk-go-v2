// Copyright (c) ZStack.io, Inc.

package param

// AttachVmNicToVmDetailParam AttachVmNicToVm详细参数
type AttachVmNicToVmDetailParam struct {
	rest string `json:"vmNicUuid" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// AttachVmNicToVmParam AttachVmNicToVm请求参数
type AttachVmNicToVmParam struct {
	BaseParam
	Params AttachVmNicToVmDetailParam `json:"params"` // 详细参数
}

