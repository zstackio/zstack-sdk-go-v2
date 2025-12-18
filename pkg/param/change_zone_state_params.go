// Copyright (c) ZStack.io, Inc.

package param

// ChangeZoneStateDetailParam ChangeZoneState detail param
type ChangeZoneStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeZoneStateParam ChangeZoneState request param
type ChangeZoneStateParam struct {
	BaseParam
	Params ChangeZoneStateDetailParam `json:"params"`
}
