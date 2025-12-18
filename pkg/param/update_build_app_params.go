// Copyright (c) ZStack.io, Inc.

package param

// UpdateBuildAppDetailParam UpdateBuildApp detail param
type UpdateBuildAppDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Version string `json:"version,omitempty"`
}

// UpdateBuildAppParam UpdateBuildApp request param
type UpdateBuildAppParam struct {
	BaseParam
	Params UpdateBuildAppDetailParam `json:"params"`
}
