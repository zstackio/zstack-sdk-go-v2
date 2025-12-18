// Copyright (c) ZStack.io, Inc.

package param

// DeleteVmCdRomDetailParam DeleteVmCdRom详细参数
type DeleteVmCdRomDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeleteVmCdRomParam DeleteVmCdRom请求参数
type DeleteVmCdRomParam struct {
	BaseParam
	Params DeleteVmCdRomDetailParam `json:"params"` // 详细参数
}

