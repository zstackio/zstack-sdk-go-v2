// Copyright (c) ZStack.io, Inc.

package param

// CreateIAM2ProjectTemplateFromProjectDetailParam CreateIAM2ProjectTemplateFromProject详细参数
type CreateIAM2ProjectTemplateFromProjectDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"projectUuid" validate:"required"` // 必填
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateIAM2ProjectTemplateFromProjectParam CreateIAM2ProjectTemplateFromProject请求参数
type CreateIAM2ProjectTemplateFromProjectParam struct {
	BaseParam
	Params CreateIAM2ProjectTemplateFromProjectDetailParam `json:"params"` // 详细参数
}

