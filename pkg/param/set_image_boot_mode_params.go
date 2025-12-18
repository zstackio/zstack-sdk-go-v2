// Copyright (c) ZStack.io, Inc.

package param

// SetImageBootModeDetailParam SetImageBootMode详细参数
type SetImageBootModeDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"bootMode" validate:"required"` // 必填
}

// SetImageBootModeParam SetImageBootMode请求参数
type SetImageBootModeParam struct {
	BaseParam
	Params SetImageBootModeDetailParam `json:"params"` // 详细参数
}

