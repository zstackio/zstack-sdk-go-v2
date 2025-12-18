// Copyright (c) ZStack.io, Inc.

package param

// GetVmsCapabilitiesDetailParam GetVmsCapabilities detail param
type GetVmsCapabilitiesDetailParam struct {
	VmUuids []string `json:"vmUuids" validate:"required"`
}

// GetVmsCapabilitiesParam GetVmsCapabilities request param
type GetVmsCapabilitiesParam struct {
	BaseParam
	Params GetVmsCapabilitiesDetailParam `json:"params"`
}
