// Copyright (c) ZStack.io, Inc.

package param

// SetVmHostnameDetailParam SetVmHostname detail param
type SetVmHostnameDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Hostname string `json:"hostname" validate:"required"`
}

// SetVmHostnameParam SetVmHostname request param
type SetVmHostnameParam struct {
	BaseParam
	Params SetVmHostnameDetailParam `json:"params"`
}
