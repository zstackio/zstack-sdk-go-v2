// Copyright (c) ZStack.io, Inc.

package param

// ChangeAppBuildSystemStateDetailParam ChangeAppBuildSystemState detail param
type ChangeAppBuildSystemStateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	StateEvent string `json:"stateEvent" validate:"required"`
}

// ChangeAppBuildSystemStateParam ChangeAppBuildSystemState request param
type ChangeAppBuildSystemStateParam struct {
	BaseParam
	Params ChangeAppBuildSystemStateDetailParam `json:"params"`
}
