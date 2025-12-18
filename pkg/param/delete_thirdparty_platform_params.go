// Copyright (c) ZStack.io, Inc.

package param

// DeleteThirdpartyPlatformDetailParam DeleteThirdpartyPlatform detail param
type DeleteThirdpartyPlatformDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteThirdpartyPlatformParam DeleteThirdpartyPlatform request param
type DeleteThirdpartyPlatformParam struct {
	BaseParam
	Params DeleteThirdpartyPlatformDetailParam `json:"params"`
}
