// Copyright (c) ZStack.io, Inc.

package param

// UpdateMdevDeviceSpecDetailParam UpdateMdevDeviceSpec detail param
type UpdateMdevDeviceSpecDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
}

// UpdateMdevDeviceSpecParam UpdateMdevDeviceSpec request param
type UpdateMdevDeviceSpecParam struct {
	BaseParam
	Params UpdateMdevDeviceSpecDetailParam `json:"params"`
}
