// Copyright (c) ZStack.io, Inc.

package param

// UpdateHostDetailParam UpdateHost detail param
type UpdateHostDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
}

// UpdateHostParam UpdateHost request param
type UpdateHostParam struct {
	BaseParam
	Params UpdateHostDetailParam `json:"params"`
}
