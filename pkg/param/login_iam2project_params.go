// Copyright (c) ZStack.io, Inc.

package param

// LoginIAM2ProjectDetailParam LoginIAM2Project详细参数
type LoginIAM2ProjectDetailParam struct {
	rest string `json:"projectName" validate:"required"` // 必填
	rest map[string]string `json:"clientInfo,omitempty"`
}

// LoginIAM2ProjectParam LoginIAM2Project请求参数
type LoginIAM2ProjectParam struct {
	BaseParam
	Params LoginIAM2ProjectDetailParam `json:"params"` // 详细参数
}

