// Copyright (c) ZStack.io, Inc.

package param

// SetVolumeIoThreadPinDetailParam SetVolumeIoThreadPin detail param
type SetVolumeIoThreadPinDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	VmUuid string `json:"vmUuid" validate:"required"`
	Pin string `json:"pin" validate:"required"`
	IoThreadId int `json:"ioThreadId" validate:"required"`
}

// SetVolumeIoThreadPinParam SetVolumeIoThreadPin request param
type SetVolumeIoThreadPinParam struct {
	BaseParam
	Params SetVolumeIoThreadPinDetailParam `json:"params"`
}
