// Copyright (c) ZStack.io, Inc.

package param

// DeletePreconfigurationTemplateDetailParam DeletePreconfigurationTemplate detail param
type DeletePreconfigurationTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeletePreconfigurationTemplateParam DeletePreconfigurationTemplate request param
type DeletePreconfigurationTemplateParam struct {
	BaseParam
	Params DeletePreconfigurationTemplateDetailParam `json:"params"`
}
