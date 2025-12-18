// Copyright (c) ZStack.io, Inc.

package param

// GetHostMultipathTopologyDetailParam GetHostMultipathTopology detail param
type GetHostMultipathTopologyDetailParam struct {
	HostUuid string `json:"hostUuid" validate:"required"`
	DiskUuids []string `json:"diskUuids" validate:"required"`
}

// GetHostMultipathTopologyParam GetHostMultipathTopology request param
type GetHostMultipathTopologyParam struct {
	BaseParam
	Params GetHostMultipathTopologyDetailParam `json:"params"`
}
