// Copyright (c) ZStack.io, Inc.

package param

// UpdateBareMetal2ChassisOfferingDetailParam UpdateBareMetal2ChassisOffering detail param
type UpdateBareMetal2ChassisOfferingDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateBareMetal2ChassisOfferingParam UpdateBareMetal2ChassisOffering request param
type UpdateBareMetal2ChassisOfferingParam struct {
	BaseParam
	Params UpdateBareMetal2ChassisOfferingDetailParam `json:"params"`
}
