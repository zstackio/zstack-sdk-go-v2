// Copyright (c) ZStack.io, Inc.

package param

// SetVmQxlMemoryDetailParam SetVmQxlMemory详细参数
type SetVmQxlMemoryDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int `json:"ram,omitempty"`
	rest int `json:"vram,omitempty"`
	rest int `json:"vgamem,omitempty"`
}

// SetVmQxlMemoryParam SetVmQxlMemory请求参数
type SetVmQxlMemoryParam struct {
	BaseParam
	Params SetVmQxlMemoryDetailParam `json:"params"` // 详细参数
}

