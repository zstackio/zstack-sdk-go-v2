// Copyright (c) ZStack.io, Inc.

package param

// UpdateSNSApplicationPlatformDetailParam UpdateSNSApplicationPlatform detail param
type UpdateSNSApplicationPlatformDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateSNSApplicationPlatformParam UpdateSNSApplicationPlatform request param
type UpdateSNSApplicationPlatformParam struct {
	BaseParam
	Params UpdateSNSApplicationPlatformDetailParam `json:"params"`
}
