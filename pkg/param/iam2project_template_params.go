// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateIAM2ProjectTemplateParamDetail CreateIAM2ProjectTemplate detail param
type CreateIAM2ProjectTemplateParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Attributes []AttributeParam `json:"attributes,omitempty"`
	Quota map[string]int64 `json:"quota,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIAM2ProjectTemplateParam CreateIAM2ProjectTemplate request param
type CreateIAM2ProjectTemplateParam struct {
	BaseParam
	Params CreateIAM2ProjectTemplateParamDetail `json:"params"`
}
// UpdateIAM2ProjectTemplateParamDetail UpdateIAM2ProjectTemplate detail param
type UpdateIAM2ProjectTemplateParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Attributes []AttributeParam `json:"attributes,omitempty"`
	Quota map[string]int64 `json:"quota,omitempty"`
}

// UpdateIAM2ProjectTemplateParam UpdateIAM2ProjectTemplate request param
type UpdateIAM2ProjectTemplateParam struct {
	BaseParam
	Params UpdateIAM2ProjectTemplateParamDetail `json:"updateIAM2ProjectTemplate"`
}
// DeleteIAM2ProjectTemplateParamDetail DeleteIAM2ProjectTemplate detail param
type DeleteIAM2ProjectTemplateParamDetail struct {
}

// DeleteIAM2ProjectTemplateParam DeleteIAM2ProjectTemplate request param
type DeleteIAM2ProjectTemplateParam struct {
	BaseParam
	Params DeleteIAM2ProjectTemplateParamDetail `json:"deleteIAM2ProjectTemplate"`
}
