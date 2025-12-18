// Copyright (c) ZStack.io, Inc.

package param

// ChangeSchedulerStateDetailParam ChangeSchedulerState detail param
type ChangeSchedulerStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeSchedulerStateParam ChangeSchedulerState request param
type ChangeSchedulerStateParam struct {
	BaseParam
	Params ChangeSchedulerStateDetailParam `json:"params"`
}
