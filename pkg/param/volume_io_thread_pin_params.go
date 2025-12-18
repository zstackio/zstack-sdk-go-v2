// Copyright (c) ZStack.io, Inc.

package param

// GetVolumeIoThreadPinDetailParam GetVolumeIoThreadPin详细参数
type GetVolumeIoThreadPinDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVolumeIoThreadPinParam GetVolumeIoThreadPin请求参数
type GetVolumeIoThreadPinParam struct {
	BaseParam
	Params GetVolumeIoThreadPinDetailParam `json:"params"` // 详细参数
}

