// Copyright (c) ZStack.io, Inc.

package param

// GetVmInstanceFirstBootDeviceDetailParam GetVmInstanceFirstBootDevice详细参数
type GetVmInstanceFirstBootDeviceDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVmInstanceFirstBootDeviceParam GetVmInstanceFirstBootDevice请求参数
type GetVmInstanceFirstBootDeviceParam struct {
	BaseParam
	Params GetVmInstanceFirstBootDeviceDetailParam `json:"params"` // 详细参数
}

