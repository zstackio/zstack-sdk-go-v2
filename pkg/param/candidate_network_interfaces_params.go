// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateNetworkInterfacesDetailParam GetCandidateNetworkInterfaces详细参数
type GetCandidateNetworkInterfacesDetailParam struct {
	rest []string `json:"hostUuids" validate:"required"` // 必填
	rest string `json:"interfaceType,omitempty"`
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetCandidateNetworkInterfacesParam GetCandidateNetworkInterfaces请求参数
type GetCandidateNetworkInterfacesParam struct {
	BaseParam
	Params GetCandidateNetworkInterfacesDetailParam `json:"params"` // 详细参数
}

