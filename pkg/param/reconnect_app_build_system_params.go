// Copyright (c) ZStack.io, Inc.

package param

// ReconnectAppBuildSystemDetailParam ReconnectAppBuildSystem detail param
type ReconnectAppBuildSystemDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// ReconnectAppBuildSystemParam ReconnectAppBuildSystem request param
type ReconnectAppBuildSystemParam struct {
	BaseParam
	Params ReconnectAppBuildSystemDetailParam `json:"params"`
}
