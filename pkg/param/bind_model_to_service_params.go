// Copyright (c) ZStack.io, Inc.

package param

// BindModelToServiceDetailParam BindModelToService详细参数
type BindModelToServiceDetailParam struct {
	rest string `json:"modelUuid" validate:"required"` // 必填
	rest string `json:"modelServiceUuid" validate:"required"` // 必填
}

// BindModelToServiceParam BindModelToService请求参数
type BindModelToServiceParam struct {
	BaseParam
	Params BindModelToServiceDetailParam `json:"params"` // 详细参数
}

