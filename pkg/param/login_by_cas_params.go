// Copyright (c) ZStack.io, Inc.

package param

// LoginByCasDetailParam LoginByCas详细参数
type LoginByCasDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
	rest map[string]string `json:"clientInfo,omitempty"`
}

// LoginByCasParam LoginByCas请求参数
type LoginByCasParam struct {
	BaseParam
	Params LoginByCasDetailParam `json:"params"` // 详细参数
}

