// Copyright (c) ZStack.io, Inc.

package param

// CreateUserTagDetailParam CreateUserTag详细参数
type CreateUserTagDetailParam struct {
	rest string `json:"resourceType" validate:"required"` // 必填
	rest string `json:"resourceUuid" validate:"required"` // 必填
	rest string `json:"tag" validate:"required"` // 必填
}

// CreateUserTagParam CreateUserTag请求参数
type CreateUserTagParam struct {
	BaseParam
	Params CreateUserTagDetailParam `json:"params"` // 详细参数
}

