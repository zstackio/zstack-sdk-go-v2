// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2ProjectContainerImagesDetailParam GetIAM2ProjectContainerImages详细参数
type GetIAM2ProjectContainerImagesDetailParam struct {
	rest string `json:"projectId" validate:"required"` // 必填
	rest string `json:"repositoryId" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetIAM2ProjectContainerImagesParam GetIAM2ProjectContainerImages请求参数
type GetIAM2ProjectContainerImagesParam struct {
	BaseParam
	Params GetIAM2ProjectContainerImagesDetailParam `json:"params"` // 详细参数
}

