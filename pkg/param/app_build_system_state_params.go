// Copyright (c) ZStack.io, Inc.

package param

// ChangeAppBuildSystemStateDetailParam ChangeAppBuildSystemState详细参数
type ChangeAppBuildSystemStateDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"stateEvent" validate:"required"` // 必填
}

// ChangeAppBuildSystemStateParam ChangeAppBuildSystemState请求参数
type ChangeAppBuildSystemStateParam struct {
	BaseParam
	Params ChangeAppBuildSystemStateDetailParam `json:"params"` // 详细参数
}

