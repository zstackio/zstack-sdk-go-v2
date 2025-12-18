// Copyright (c) ZStack.io, Inc.

package param

// DeleteIAM2ProjectDetailParam DeleteIAM2Project详细参数
type DeleteIAM2ProjectDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// DeleteIAM2ProjectParam DeleteIAM2Project请求参数
type DeleteIAM2ProjectParam struct {
	BaseParam
	Params DeleteIAM2ProjectDetailParam `json:"params"` // 详细参数
}

