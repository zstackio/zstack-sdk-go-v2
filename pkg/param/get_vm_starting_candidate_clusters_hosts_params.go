// Copyright (c) ZStack.io, Inc.

package param

// GetVmStartingCandidateClustersHostsDetailParam GetVmStartingCandidateClustersHosts detail param
type GetVmStartingCandidateClustersHostsDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// GetVmStartingCandidateClustersHostsParam GetVmStartingCandidateClustersHosts request param
type GetVmStartingCandidateClustersHostsParam struct {
	BaseParam
	Params GetVmStartingCandidateClustersHostsDetailParam `json:"params"`
}
