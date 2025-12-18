// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2ProjectFromTemplateDetailParam CreateIAM2ProjectFromTemplate详细参数
type CreateIAM2ProjectFromTemplateDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"templateUuid" validate:"required"` // 必填
	rest []string `json:"roleUuids,omitempty"`
	rest string `json:"organizationUuid,omitempty"`
	rest []string `json:"resourceTemplates,omitempty"`
	rest string `json:"linkAccountUuid,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateIAM2ProjectFromTemplateParam CreateIAM2ProjectFromTemplate请求参数
type CreateIAM2ProjectFromTemplateParam struct {
	BaseParam
	Params CreateIAM2ProjectFromTemplateDetailParam `json:"params"` // 详细参数
}

