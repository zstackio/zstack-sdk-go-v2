// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2ProjectContainerImageTagsDetailParam GetIAM2ProjectContainerImageTags详细参数
type GetIAM2ProjectContainerImageTagsDetailParam struct {
	rest string `json:"projectId" validate:"required"` // 必填
	rest string `json:"repositoryId" validate:"required"` // 必填
	rest string `json:"imageName" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetIAM2ProjectContainerImageTagsParam GetIAM2ProjectContainerImageTags请求参数
type GetIAM2ProjectContainerImageTagsParam struct {
	BaseParam
	Params GetIAM2ProjectContainerImageTagsDetailParam `json:"params"` // 详细参数
}

