// Copyright (c) ZStack.io, Inc.

package param

// TokenIntrospectionDetailParam TokenIntrospection detail param
type TokenIntrospectionDetailParam struct {
	Token string `json:"token" validate:"required"`
	TokenType string `json:"tokenType" validate:"required"`
}

// TokenIntrospectionParam TokenIntrospection request param
type TokenIntrospectionParam struct {
	BaseParam
	Params TokenIntrospectionDetailParam `json:"params"`
}
