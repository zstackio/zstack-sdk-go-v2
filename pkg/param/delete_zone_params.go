// Copyright (c) ZStack.io, Inc.

package param

// DeleteZoneDetailParam DeleteZone detail param
type DeleteZoneDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteZoneParam DeleteZone request param
type DeleteZoneParam struct {
	BaseParam
	Params DeleteZoneDetailParam `json:"params"`
}
