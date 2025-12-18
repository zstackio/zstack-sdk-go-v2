// Copyright (c) ZStack.io, Inc.

package param

// GetVmvNUMATopologyDetailParam GetVmvNUMATopology detail param
type GetVmvNUMATopologyDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmvNUMATopologyParam GetVmvNUMATopology request param
type GetVmvNUMATopologyParam struct {
	BaseParam
	Params GetVmvNUMATopologyDetailParam `json:"params"`
}
