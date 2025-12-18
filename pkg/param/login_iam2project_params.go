// Copyright (c) ZStack.io, Inc.

package param

// LoginIAM2ProjectDetailParam LoginIAM2Project detail param
type LoginIAM2ProjectDetailParam struct {
	ProjectName string `json:"projectName" validate:"required"`
	ClientInfo map[string]string `json:"clientInfo,omitempty"`
}

// LoginIAM2ProjectParam LoginIAM2Project request param
type LoginIAM2ProjectParam struct {
	BaseParam
	Params LoginIAM2ProjectDetailParam `json:"params"`
}
