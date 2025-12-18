// Copyright (c) ZStack.io, Inc.

package param

// DeleteBuildAppDetailParam DeleteBuildApp detail param
type DeleteBuildAppDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteBuildAppParam DeleteBuildApp request param
type DeleteBuildAppParam struct {
	BaseParam
	Params DeleteBuildAppDetailParam `json:"params"`
}
