// Copyright (c) ZStack.io, Inc.

package param

// GetVolumeIoThreadPinDetailParam GetVolumeIoThreadPin detail param
type GetVolumeIoThreadPinDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVolumeIoThreadPinParam GetVolumeIoThreadPin request param
type GetVolumeIoThreadPinParam struct {
	BaseParam
	Params GetVolumeIoThreadPinDetailParam `json:"params"`
}
