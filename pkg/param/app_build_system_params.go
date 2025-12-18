// Copyright (c) ZStack.io, Inc.

package param

// DeleteAppBuildSystemDetailParam DeleteAppBuildSystem详细参数
type DeleteAppBuildSystemDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteAppBuildSystemParam DeleteAppBuildSystem请求参数
type DeleteAppBuildSystemParam struct {
	BaseParam
	Params DeleteAppBuildSystemDetailParam `json:"params"` // 详细参数
}

