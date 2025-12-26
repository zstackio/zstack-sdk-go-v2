// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2ProjectTemplateDetailParam CreateIAM2ProjectTemplate detail param
type CreateIAM2ProjectTemplateDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Attributes []AttributeParam `json:"attributes,omitempty"`
	Quota map[string]int64 `json:"quota,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2ProjectTemplateParam CreateIAM2ProjectTemplate request param
type CreateIAM2ProjectTemplateParam struct {
	BaseParam
	Params CreateIAM2ProjectTemplateDetailParam `json:"params"`
}
