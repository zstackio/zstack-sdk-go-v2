// Copyright (c) ZStack.io, Inc.

package param

// UpdateMdevDeviceDetailParam UpdateMdevDevice detail param
type UpdateMdevDeviceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
}

// UpdateMdevDeviceParam UpdateMdevDevice request param
type UpdateMdevDeviceParam struct {
	BaseParam
	Params UpdateMdevDeviceDetailParam `json:"params"`
}
