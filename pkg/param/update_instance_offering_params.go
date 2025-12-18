// Copyright (c) ZStack.io, Inc.

package param

// UpdateInstanceOfferingDetailParam UpdateInstanceOffering detail param
type UpdateInstanceOfferingDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AllocatorStrategy string `json:"allocatorStrategy,omitempty"`
}

// UpdateInstanceOfferingParam UpdateInstanceOffering request param
type UpdateInstanceOfferingParam struct {
	BaseParam
	Params UpdateInstanceOfferingDetailParam `json:"params"`
}
