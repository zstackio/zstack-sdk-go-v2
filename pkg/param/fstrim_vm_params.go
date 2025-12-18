// Copyright (c) ZStack.io, Inc.

package param

// FstrimVmDetailParam FstrimVm详细参数
type FstrimVmDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// FstrimVmParam FstrimVm请求参数
type FstrimVmParam struct {
	BaseParam
	Params FstrimVmDetailParam `json:"params"` // 详细参数
}

