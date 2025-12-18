// Copyright (c) ZStack.io, Inc.

package param

// GetVmvNUMATopologyDetailParam GetVmvNUMATopology详细参数
type GetVmvNUMATopologyDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVmvNUMATopologyParam GetVmvNUMATopology请求参数
type GetVmvNUMATopologyParam struct {
	BaseParam
	Params GetVmvNUMATopologyDetailParam `json:"params"` // 详细参数
}

