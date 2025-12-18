// Copyright (c) ZStack.io, Inc.

package param

// UpdateGlobalConfigDetailParam UpdateGlobalConfig detail param
type UpdateGlobalConfigDetailParam struct {
	Category string `json:"category" validate:"required"`
	Name string `json:"name" validate:"required"`
	Value string `json:"value,omitempty"`
}

// UpdateGlobalConfigParam UpdateGlobalConfig request param
type UpdateGlobalConfigParam struct {
	BaseParam
	Params UpdateGlobalConfigDetailParam `json:"params"`
}
