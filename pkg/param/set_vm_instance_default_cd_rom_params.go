// Copyright (c) ZStack.io, Inc.

package param

// SetVmInstanceDefaultCdRomDetailParam SetVmInstanceDefaultCdRom详细参数
type SetVmInstanceDefaultCdRomDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// SetVmInstanceDefaultCdRomParam SetVmInstanceDefaultCdRom请求参数
type SetVmInstanceDefaultCdRomParam struct {
	BaseParam
	Params SetVmInstanceDefaultCdRomDetailParam `json:"params"` // 详细参数
}

