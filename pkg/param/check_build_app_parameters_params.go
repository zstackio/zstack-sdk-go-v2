// Copyright (c) ZStack.io, Inc.

package param

// CheckBuildAppParametersDetailParam CheckBuildAppParameters详细参数
type CheckBuildAppParametersDetailParam struct {
	rest string `json:"type,omitempty"`
	rest string `json:"uuid" validate:"required"` // 必填
}

// CheckBuildAppParametersParam CheckBuildAppParameters请求参数
type CheckBuildAppParametersParam struct {
	BaseParam
	Params CheckBuildAppParametersDetailParam `json:"params"` // 详细参数
}

