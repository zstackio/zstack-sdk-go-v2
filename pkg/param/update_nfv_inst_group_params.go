// Copyright (c) ZStack.io, Inc.

package param

// UpdateNfvInstGroupDetailParam UpdateNfvInstGroup detail param
type UpdateNfvInstGroupDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateNfvInstGroupParam UpdateNfvInstGroup request param
type UpdateNfvInstGroupParam struct {
	BaseParam
	Params UpdateNfvInstGroupDetailParam `json:"params"`
}
