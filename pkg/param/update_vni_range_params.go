// Copyright (c) ZStack.io, Inc.

package param

// UpdateVniRangeDetailParam UpdateVniRange detail param
type UpdateVniRangeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name" validate:"required"`
}

// UpdateVniRangeParam UpdateVniRange request param
type UpdateVniRangeParam struct {
	BaseParam
	Params UpdateVniRangeDetailParam `json:"params"`
}
