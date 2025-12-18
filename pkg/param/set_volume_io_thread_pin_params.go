// Copyright (c) ZStack.io, Inc.

package param

// SetVolumeIoThreadPinDetailParam SetVolumeIoThreadPin详细参数
type SetVolumeIoThreadPinDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"vmUuid" validate:"required"` // 必填
	rest string `json:"pin" validate:"required"` // 必填
	rest int `json:"ioThreadId" validate:"required"` // 必填
}

// SetVolumeIoThreadPinParam SetVolumeIoThreadPin请求参数
type SetVolumeIoThreadPinParam struct {
	BaseParam
	Params SetVolumeIoThreadPinDetailParam `json:"params"` // 详细参数
}

