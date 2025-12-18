// Copyright (c) ZStack.io, Inc.

package param

// UpdateVmInstanceDetailParam UpdateVmInstance detail param
type UpdateVmInstanceDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty"`
	Platform string `json:"platform,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	ReservedMemorySize int64 `json:"reservedMemorySize,omitempty"`
	GuestOsType string `json:"guestOsType,omitempty"`
}

// UpdateVmInstanceParam UpdateVmInstance request param
type UpdateVmInstanceParam struct {
	BaseParam
	Params UpdateVmInstanceDetailParam `json:"params"`
}
