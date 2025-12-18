// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2ProjectContainerImageTagsDetailParam GetIAM2ProjectContainerImageTags detail param
type GetIAM2ProjectContainerImageTagsDetailParam struct {
	ProjectId string `json:"projectId" validate:"required"`
	RepositoryId string `json:"repositoryId" validate:"required"`
	ImageName string `json:"imageName" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetIAM2ProjectContainerImageTagsParam GetIAM2ProjectContainerImageTags request param
type GetIAM2ProjectContainerImageTagsParam struct {
	BaseParam
	Params GetIAM2ProjectContainerImageTagsDetailParam `json:"params"`
}
