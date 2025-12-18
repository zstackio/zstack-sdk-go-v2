// Copyright (c) ZStack.io, Inc.

package param

// GetVmHostnameDetailParam GetVmHostname detail param
type GetVmHostnameDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmHostnameParam GetVmHostname request param
type GetVmHostnameParam struct {
	BaseParam
	Params GetVmHostnameDetailParam `json:"params"`
}
