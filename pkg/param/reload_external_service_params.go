// Copyright (c) ZStack.io, Inc.

package param

// ReloadExternalServiceDetailParam ReloadExternalService详细参数
type ReloadExternalServiceDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
}

// ReloadExternalServiceParam ReloadExternalService请求参数
type ReloadExternalServiceParam struct {
	BaseParam
	Params ReloadExternalServiceDetailParam `json:"params"` // 详细参数
}

