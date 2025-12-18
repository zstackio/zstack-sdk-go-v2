// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2ProjectRepositoryDetailParam GetIAM2ProjectRepository详细参数
type GetIAM2ProjectRepositoryDetailParam struct {
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetIAM2ProjectRepositoryParam GetIAM2ProjectRepository请求参数
type GetIAM2ProjectRepositoryParam struct {
	BaseParam
	Params GetIAM2ProjectRepositoryDetailParam `json:"params"` // 详细参数
}

