// Copyright (c) ZStack.io, Inc.

package param

// TokenIntrospectionDetailParam TokenIntrospection详细参数
type TokenIntrospectionDetailParam struct {
	rest string `json:"token" validate:"required"` // 必填
	rest string `json:"tokenType" validate:"required"` // 必填
}

// TokenIntrospectionParam TokenIntrospection请求参数
type TokenIntrospectionParam struct {
	BaseParam
	Params TokenIntrospectionDetailParam `json:"params"` // 详细参数
}

