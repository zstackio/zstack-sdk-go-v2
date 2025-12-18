// Copyright (c) ZStack.io, Inc.

package param

// AddAttributesToIAM2VirtualIDDetailParam AddAttributesToIAM2VirtualID详细参数
type AddAttributesToIAM2VirtualIDDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []interface{} `json:"attributes" validate:"required"` // 必填
}

// AddAttributesToIAM2VirtualIDParam AddAttributesToIAM2VirtualID请求参数
type AddAttributesToIAM2VirtualIDParam struct {
	BaseParam
	Params AddAttributesToIAM2VirtualIDDetailParam `json:"params"` // 详细参数
}

