// Copyright (c) ZStack.io, Inc.

package param

// GetVmStartingCandidateClustersHostsDetailParam GetVmStartingCandidateClustersHosts详细参数
type GetVmStartingCandidateClustersHostsDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetVmStartingCandidateClustersHostsParam GetVmStartingCandidateClustersHosts请求参数
type GetVmStartingCandidateClustersHostsParam struct {
	BaseParam
	Params GetVmStartingCandidateClustersHostsDetailParam `json:"params"` // 详细参数
}

