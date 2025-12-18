// Copyright (c) ZStack.io, Inc.

package param

// UpdateIAM2ProjectTemplateDetailParam UpdateIAM2ProjectTemplate detail param
type UpdateIAM2ProjectTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Attributes []interface{} `json:"attributes,omitempty"`
	Quota map[string]int64 `json:"quota,omitempty"`
}

// UpdateIAM2ProjectTemplateParam UpdateIAM2ProjectTemplate request param
type UpdateIAM2ProjectTemplateParam struct {
	BaseParam
	Params UpdateIAM2ProjectTemplateDetailParam `json:"params"`
}
