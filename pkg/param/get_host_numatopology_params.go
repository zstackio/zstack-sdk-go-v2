// Copyright (c) ZStack.io, Inc.

package param

// GetHostNUMATopologyDetailParam GetHostNUMATopology detail param
type GetHostNUMATopologyDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetHostNUMATopologyParam GetHostNUMATopology request param
type GetHostNUMATopologyParam struct {
	BaseParam
	Params GetHostNUMATopologyDetailParam `json:"params"`
}
