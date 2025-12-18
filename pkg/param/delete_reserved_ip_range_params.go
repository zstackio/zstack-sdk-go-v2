// Copyright (c) ZStack.io, Inc.

package param

// DeleteReservedIpRangeDetailParam DeleteReservedIpRange detail param
type DeleteReservedIpRangeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteReservedIpRangeParam DeleteReservedIpRange request param
type DeleteReservedIpRangeParam struct {
	BaseParam
	Params DeleteReservedIpRangeDetailParam `json:"params"`
}
