// Copyright (c) ZStack.io, Inc.

package param

// DecodeStackTemplateDetailParam DecodeStackTemplate detail param
type DecodeStackTemplateDetailParam struct {
	Type string `json:"type,omitempty"`
	TemplateContent string `json:"templateContent,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Parameters string `json:"parameters,omitempty"`
	Preparameters string `json:"preparameters,omitempty"`
}

// DecodeStackTemplateParam DecodeStackTemplate request param
type DecodeStackTemplateParam struct {
	BaseParam
	Params DecodeStackTemplateDetailParam `json:"params"`
}
