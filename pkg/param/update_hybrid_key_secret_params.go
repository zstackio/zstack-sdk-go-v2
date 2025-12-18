// Copyright (c) ZStack.io, Inc.

package param

// UpdateHybridKeySecretDetailParam UpdateHybridKeySecret detail param
type UpdateHybridKeySecretDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateHybridKeySecretParam UpdateHybridKeySecret request param
type UpdateHybridKeySecretParam struct {
	BaseParam
	Params UpdateHybridKeySecretDetailParam `json:"params"`
}
