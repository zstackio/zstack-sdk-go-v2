// Copyright (c) ZStack.io, Inc.

package param

// AddReservedIpRangeDetailParam AddReservedIpRange detail param
type AddReservedIpRangeDetailParam struct {
	L3NetworkUuid string `json:"l3NetworkUuid" validate:"required"`
	StartIp string `json:"startIp" validate:"required"`
	EndIp string `json:"endIp" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddReservedIpRangeParam AddReservedIpRange request param
type AddReservedIpRangeParam struct {
	BaseParam
	Params AddReservedIpRangeDetailParam `json:"params"`
}
