// Copyright (c) ZStack.io, Inc.

package param

// GetCandidateNetworkInterfacesDetailParam GetCandidateNetworkInterfaces detail param
type GetCandidateNetworkInterfacesDetailParam struct {
	HostUuids []string `json:"hostUuids" validate:"required"`
	InterfaceType string `json:"interfaceType,omitempty"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetCandidateNetworkInterfacesParam GetCandidateNetworkInterfaces request param
type GetCandidateNetworkInterfacesParam struct {
	BaseParam
	Params GetCandidateNetworkInterfacesDetailParam `json:"params"`
}
