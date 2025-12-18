// Copyright (c) ZStack.io, Inc.

package param

// GetVolumeCapabilitiesDetailParam GetVolumeCapabilities detail param
type GetVolumeCapabilitiesDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVolumeCapabilitiesParam GetVolumeCapabilities request param
type GetVolumeCapabilitiesParam struct {
	BaseParam
	Params GetVolumeCapabilitiesDetailParam `json:"params"`
}
