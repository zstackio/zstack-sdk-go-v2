// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateInterfaceVlanIdsDetailParam GetCandidateInterfaceVlanIds详细参数
type GetCandidateInterfaceVlanIdsDetailParam struct {
	rest []string `json:"interfaceUuids" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetCandidateInterfaceVlanIdsParam GetCandidateInterfaceVlanIds请求参数
type GetCandidateInterfaceVlanIdsParam struct {
	BaseParam
	Params GetCandidateInterfaceVlanIdsDetailParam `json:"params"` // 详细参数
}

