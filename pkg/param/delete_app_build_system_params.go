// Copyright (c) ZStack.io, Inc.

package param

// DeleteAppBuildSystemDetailParam DeleteAppBuildSystem detail param
type DeleteAppBuildSystemDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteAppBuildSystemParam DeleteAppBuildSystem request param
type DeleteAppBuildSystemParam struct {
	BaseParam
	Params DeleteAppBuildSystemDetailParam `json:"params"`
}
