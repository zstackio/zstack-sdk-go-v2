// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateMiniHostsDetailParam GetCandidateMiniHosts detail param
type GetCandidateMiniHostsDetailParam struct {
	Local bool `json:"local,omitempty"`
	Configure bool `json:"configure,omitempty"`
}

// GetCandidateMiniHostsParam GetCandidateMiniHosts request param
type GetCandidateMiniHostsParam struct {
	BaseParam
	Params GetCandidateMiniHostsDetailParam `json:"params"`
}
