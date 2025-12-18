// Copyright (c) ZStack.io, Inc.

package param

// UpdateEcsVSwitchDetailParam UpdateEcsVSwitch detail param
type UpdateEcsVSwitchDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateEcsVSwitchParam UpdateEcsVSwitch request param
type UpdateEcsVSwitchParam struct {
	BaseParam
	Params UpdateEcsVSwitchDetailParam `json:"params"`
}
