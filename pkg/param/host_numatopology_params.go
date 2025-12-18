// Copyright (c) ZStack.io, Inc.

package param

// GetHostNUMATopologyDetailParam GetHostNUMATopology详细参数
type GetHostNUMATopologyDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetHostNUMATopologyParam GetHostNUMATopology请求参数
type GetHostNUMATopologyParam struct {
	BaseParam
	Params GetHostNUMATopologyDetailParam `json:"params"` // 详细参数
}

