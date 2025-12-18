// Copyright (c) ZStack.io, Inc.

package param

// DeleteStackTemplateDetailParam DeleteStackTemplate detail param
type DeleteStackTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteStackTemplateParam DeleteStackTemplate request param
type DeleteStackTemplateParam struct {
	BaseParam
	Params DeleteStackTemplateDetailParam `json:"params"`
}
