// Copyright (c) ZStack.io, Inc.

package param

// RemovePciDeviceSpecFromVmInstanceDetailParam RemovePciDeviceSpecFromVmInstance详细参数
type RemovePciDeviceSpecFromVmInstanceDetailParam struct {
	rest string `json:"pciSpecUuid" validate:"required"` // 必填
	rest string `json:"vmInstanceUuid" validate:"required"` // 必填
}

// RemovePciDeviceSpecFromVmInstanceParam RemovePciDeviceSpecFromVmInstance请求参数
type RemovePciDeviceSpecFromVmInstanceParam struct {
	BaseParam
	Params RemovePciDeviceSpecFromVmInstanceDetailParam `json:"params"` // 详细参数
}

