// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateMiniHostsDetailParam GetCandidateMiniHosts详细参数
type GetCandidateMiniHostsDetailParam struct {
	rest bool `json:"local,omitempty"`
	rest bool `json:"configure,omitempty"`
}

// GetCandidateMiniHostsParam GetCandidateMiniHosts请求参数
type GetCandidateMiniHostsParam struct {
	BaseParam
	Params GetCandidateMiniHostsDetailParam `json:"params"` // 详细参数
}

