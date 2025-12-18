// Copyright (c) ZStack.io, Inc.

package param

// SetVmNumaDetailParam SetVmNuma detail param
type SetVmNumaDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Enable bool `json:"enable" validate:"required"`
}

// SetVmNumaParam SetVmNuma request param
type SetVmNumaParam struct {
	BaseParam
	Params SetVmNumaDetailParam `json:"params"`
}
