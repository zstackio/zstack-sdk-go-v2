// Copyright (c) ZStack.io, Inc.

package param

// UnbindModelFromServiceDetailParam UnbindModelFromService详细参数
type UnbindModelFromServiceDetailParam struct {
	rest string `json:"modelUuid" validate:"required"` // 必填
	rest string `json:"modelServiceUuid" validate:"required"` // 必填
}

// UnbindModelFromServiceParam UnbindModelFromService请求参数
type UnbindModelFromServiceParam struct {
	BaseParam
	Params UnbindModelFromServiceDetailParam `json:"params"` // 详细参数
}

