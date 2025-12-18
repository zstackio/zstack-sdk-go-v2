// Copyright (c) ZStack.io, Inc.

package param

// GetVmCapabilitiesDetailParam GetVmCapabilities detail param
type GetVmCapabilitiesDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmCapabilitiesParam GetVmCapabilities request param
type GetVmCapabilitiesParam struct {
	BaseParam
	Params GetVmCapabilitiesDetailParam `json:"params"`
}
