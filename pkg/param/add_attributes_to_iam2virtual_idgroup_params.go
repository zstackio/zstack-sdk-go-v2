// Copyright (c) ZStack.io, Inc.

package param

// AddAttributesToIAM2VirtualIDGroupDetailParam AddAttributesToIAM2VirtualIDGroup详细参数
type AddAttributesToIAM2VirtualIDGroupDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []interface{} `json:"attributes" validate:"required"` // 必填
}

// AddAttributesToIAM2VirtualIDGroupParam AddAttributesToIAM2VirtualIDGroup请求参数
type AddAttributesToIAM2VirtualIDGroupParam struct {
	BaseParam
	Params AddAttributesToIAM2VirtualIDGroupDetailParam `json:"params"` // 详细参数
}

