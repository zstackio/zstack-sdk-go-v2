// Copyright (c) ZStack.io, Inc.

package param

// DeleteSNSApplicationPlatformDetailParam DeleteSNSApplicationPlatform detail param
type DeleteSNSApplicationPlatformDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteSNSApplicationPlatformParam DeleteSNSApplicationPlatform request param
type DeleteSNSApplicationPlatformParam struct {
	BaseParam
	Params DeleteSNSApplicationPlatformDetailParam `json:"params"`
}
