// Copyright (c) ZStack.io, Inc.

package param

// DeleteIpRangeDetailParam DeleteIpRange detail param
type DeleteIpRangeDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteIpRangeParam DeleteIpRange request param
type DeleteIpRangeParam struct {
	BaseParam
	Params DeleteIpRangeDetailParam `json:"params"`
}
