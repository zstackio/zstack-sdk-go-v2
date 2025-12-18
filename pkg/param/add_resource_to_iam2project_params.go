// Copyright (c) ZStack.io, Inc.

package param

// AddResourceToIAM2ProjectDetailParam AddResourceToIAM2Project detail param
type AddResourceToIAM2ProjectDetailParam struct {
	ProjectUuid string `json:"projectUuid" validate:"required"`
	ResourceTemplates []string `json:"resourceTemplates" validate:"required"`
}

// AddResourceToIAM2ProjectParam AddResourceToIAM2Project request param
type AddResourceToIAM2ProjectParam struct {
	BaseParam
	Params AddResourceToIAM2ProjectDetailParam `json:"params"`
}
