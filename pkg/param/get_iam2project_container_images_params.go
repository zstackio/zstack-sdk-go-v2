// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2ProjectContainerImagesDetailParam GetIAM2ProjectContainerImages detail param
type GetIAM2ProjectContainerImagesDetailParam struct {
	ProjectId string `json:"projectId" validate:"required"`
	RepositoryId string `json:"repositoryId" validate:"required"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetIAM2ProjectContainerImagesParam GetIAM2ProjectContainerImages request param
type GetIAM2ProjectContainerImagesParam struct {
	BaseParam
	Params GetIAM2ProjectContainerImagesDetailParam `json:"params"`
}
