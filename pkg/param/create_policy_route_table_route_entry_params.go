// Copyright (c) ZStack.io, Inc.

package param

// CreatePolicyRouteTableRouteEntryDetailParam CreatePolicyRouteTableRouteEntry detail param
type CreatePolicyRouteTableRouteEntryDetailParam struct {
	TableUuid string `json:"tableUuid" validate:"required"`
	DestinationCidr string `json:"destinationCidr" validate:"required"`
	NextHopIp string `json:"nextHopIp" validate:"required"`
	Distance int `json:"distance,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreatePolicyRouteTableRouteEntryParam CreatePolicyRouteTableRouteEntry request param
type CreatePolicyRouteTableRouteEntryParam struct {
	BaseParam
	Params CreatePolicyRouteTableRouteEntryDetailParam `json:"params"`
}
