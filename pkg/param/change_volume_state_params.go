// Copyright (c) ZStack.io, Inc.

package param

// ChangeVolumeStateDetailParam ChangeVolumeState detail param
type ChangeVolumeStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeVolumeStateParam ChangeVolumeState request param
type ChangeVolumeStateParam struct {
	BaseParam
	Params ChangeVolumeStateDetailParam `json:"params"`
}
