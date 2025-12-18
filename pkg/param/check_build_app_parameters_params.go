// Copyright (c) ZStack.io, Inc.

package param

// CheckBuildAppParametersDetailParam CheckBuildAppParameters detail param
type CheckBuildAppParametersDetailParam struct {
	Type string `json:"type,omitempty"`
	Uuid string `json:"uuid" validate:"required"`
}

// CheckBuildAppParametersParam CheckBuildAppParameters request param
type CheckBuildAppParametersParam struct {
	BaseParam
	Params CheckBuildAppParametersDetailParam `json:"params"`
}
