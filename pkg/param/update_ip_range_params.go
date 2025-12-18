// Copyright (c) ZStack.io, Inc.

package param

// UpdateIpRangeDetailParam UpdateIpRange detail param
type UpdateIpRangeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateIpRangeParam UpdateIpRange request param
type UpdateIpRangeParam struct {
	BaseParam
	Params UpdateIpRangeDetailParam `json:"params"`
}
