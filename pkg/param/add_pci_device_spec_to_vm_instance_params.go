// Copyright (c) ZStack.io, Inc.

package param

// AddPciDeviceSpecToVmInstanceDetailParam AddPciDeviceSpecToVmInstance详细参数
type AddPciDeviceSpecToVmInstanceDetailParam struct {
	rest string `json:"pciSpecUuid" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
	rest int `json:"pciDeviceNumber,omitempty"`
}

// AddPciDeviceSpecToVmInstanceParam AddPciDeviceSpecToVmInstance请求参数
type AddPciDeviceSpecToVmInstanceParam struct {
	BaseParam
	Params AddPciDeviceSpecToVmInstanceDetailParam `json:"params"` // 详细参数
}

