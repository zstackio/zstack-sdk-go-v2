// Copyright (c) ZStack.io, Inc.

package param

// SetVmQxlMemoryDetailParam SetVmQxlMemory detail param
type SetVmQxlMemoryDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Ram int `json:"ram,omitempty"`
	Vram int `json:"vram,omitempty"`
	Vgamem int `json:"vgamem,omitempty"`
}

// SetVmQxlMemoryParam SetVmQxlMemory request param
type SetVmQxlMemoryParam struct {
	BaseParam
	Params SetVmQxlMemoryDetailParam `json:"params"`
}
