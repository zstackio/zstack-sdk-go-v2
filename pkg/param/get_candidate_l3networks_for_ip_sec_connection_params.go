// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateL3NetworksForIpSecConnectionDetailParam GetCandidateL3NetworksForIpSecConnection detail param
type GetCandidateL3NetworksForIpSecConnectionDetailParam struct {
	Uuid string `json:"uuid,omitempty"`
	PublicL3Uuid string `json:"publicL3Uuid,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetCandidateL3NetworksForIpSecConnectionParam GetCandidateL3NetworksForIpSecConnection request param
type GetCandidateL3NetworksForIpSecConnectionParam struct {
	BaseParam
	Params GetCandidateL3NetworksForIpSecConnectionDetailParam `json:"params"`
}
