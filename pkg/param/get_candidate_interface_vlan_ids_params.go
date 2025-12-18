// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateInterfaceVlanIdsDetailParam GetCandidateInterfaceVlanIds detail param
type GetCandidateInterfaceVlanIdsDetailParam struct {
	InterfaceUuids []string `json:"interfaceUuids" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetCandidateInterfaceVlanIdsParam GetCandidateInterfaceVlanIds request param
type GetCandidateInterfaceVlanIdsParam struct {
	BaseParam
	Params GetCandidateInterfaceVlanIdsDetailParam `json:"params"`
}
