// Copyright (c) ZStack.io, Inc.

package param

// AddAttributesToIAM2ProjectDetailParam AddAttributesToIAM2Project详细参数
type AddAttributesToIAM2ProjectDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest []interface{} `json:"attributes" validate:"required"` // 必填
}

// AddAttributesToIAM2ProjectParam AddAttributesToIAM2Project请求参数
type AddAttributesToIAM2ProjectParam struct {
	BaseParam
	Params AddAttributesToIAM2ProjectDetailParam `json:"params"` // 详细参数
}

