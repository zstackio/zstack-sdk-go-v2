// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateL3NetworksForIpSecConnectionDetailParam GetCandidateL3NetworksForIpSecConnection详细参数
type GetCandidateL3NetworksForIpSecConnectionDetailParam struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"publicL3Uuid,omitempty"`
	rest string `json:"vipUuid,omitempty"`
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetCandidateL3NetworksForIpSecConnectionParam GetCandidateL3NetworksForIpSecConnection请求参数
type GetCandidateL3NetworksForIpSecConnectionParam struct {
	BaseParam
	Params GetCandidateL3NetworksForIpSecConnectionDetailParam `json:"params"` // 详细参数
}

