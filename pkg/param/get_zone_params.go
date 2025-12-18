// Copyright (c) ZStack.io, Inc.

package param

// GetZoneDetailParam GetZone detail param
type GetZoneDetailParam struct {
	Uuid string `json:"uuid,omitempty"`
}

// GetZoneParam GetZone request param
type GetZoneParam struct {
	BaseParam
	Params GetZoneDetailParam `json:"params"`
}
