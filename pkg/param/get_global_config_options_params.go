// Copyright (c) ZStack.io, Inc.

package param

// GetGlobalConfigOptionsDetailParam GetGlobalConfigOptions detail param
type GetGlobalConfigOptionsDetailParam struct {
	Category string `json:"category" validate:"required"`
	Name string `json:"name" validate:"required"`
}

// GetGlobalConfigOptionsParam GetGlobalConfigOptions request param
type GetGlobalConfigOptionsParam struct {
	BaseParam
	Params GetGlobalConfigOptionsDetailParam `json:"params"`
}
