// Copyright (c) ZStack.io, Inc.

package param

// GetVfPciDeviceAvailableInL2NetworkDetailParam GetVfPciDeviceAvailableInL2Network详细参数
type GetVfPciDeviceAvailableInL2NetworkDetailParam struct {
	rest []string `json:"l2NetworkUuids" validate:"required"` // 必填
}

// GetVfPciDeviceAvailableInL2NetworkParam GetVfPciDeviceAvailableInL2Network请求参数
type GetVfPciDeviceAvailableInL2NetworkParam struct {
	BaseParam
	Params GetVfPciDeviceAvailableInL2NetworkDetailParam `json:"params"` // 详细参数
}

