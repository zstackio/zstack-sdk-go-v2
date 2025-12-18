// Copyright (c) ZStack.io, Inc.

package param

// DeleteIAM2ProjectTemplateDetailParam DeleteIAM2ProjectTemplate detail param
type DeleteIAM2ProjectTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteIAM2ProjectTemplateParam DeleteIAM2ProjectTemplate request param
type DeleteIAM2ProjectTemplateParam struct {
	BaseParam
	Params DeleteIAM2ProjectTemplateDetailParam `json:"params"`
}
