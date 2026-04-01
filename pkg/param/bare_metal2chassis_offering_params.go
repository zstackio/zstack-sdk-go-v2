// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateBareMetal2ChassisOfferingParamDetail UpdateBareMetal2ChassisOffering detail param
type UpdateBareMetal2ChassisOfferingParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateBareMetal2ChassisOfferingParam UpdateBareMetal2ChassisOffering request param
type UpdateBareMetal2ChassisOfferingParam struct {
	BaseParam
	Params UpdateBareMetal2ChassisOfferingParamDetail `json:"updateBareMetal2ChassisOffering"`
}
