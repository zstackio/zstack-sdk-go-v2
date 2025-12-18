// Copyright (c) ZStack.io, Inc.

package param

// UpdateVipDetailParam UpdateVip detail param
type UpdateVipDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateVipParam UpdateVip request param
type UpdateVipParam struct {
	BaseParam
	Params UpdateVipDetailParam `json:"params"`
}
