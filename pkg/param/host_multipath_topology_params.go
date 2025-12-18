// Copyright (c) ZStack.io, Inc.

package param

// GetHostMultipathTopologyDetailParam GetHostMultipathTopology详细参数
type GetHostMultipathTopologyDetailParam struct {
	rest string `json:"hostUuid" validate:"required"` // 必填
	rest []string `json:"diskUuids" validate:"required"` // 必填
}

// GetHostMultipathTopologyParam GetHostMultipathTopology请求参数
type GetHostMultipathTopologyParam struct {
	BaseParam
	Params GetHostMultipathTopologyDetailParam `json:"params"` // 详细参数
}

