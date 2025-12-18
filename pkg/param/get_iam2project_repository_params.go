// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2ProjectRepositoryDetailParam GetIAM2ProjectRepository detail param
type GetIAM2ProjectRepositoryDetailParam struct {
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetIAM2ProjectRepositoryParam GetIAM2ProjectRepository request param
type GetIAM2ProjectRepositoryParam struct {
	BaseParam
	Params GetIAM2ProjectRepositoryDetailParam `json:"params"`
}
