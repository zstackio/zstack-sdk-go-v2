// Copyright (c) ZStack.io, Inc.

package param

// GetIAM2ProjectsOfVirtualIDDetailParam GetIAM2ProjectsOfVirtualID详细参数
type GetIAM2ProjectsOfVirtualIDDetailParam struct {
}

// GetIAM2ProjectsOfVirtualIDParam GetIAM2ProjectsOfVirtualID请求参数
type GetIAM2ProjectsOfVirtualIDParam struct {
	BaseParam
	Params GetIAM2ProjectsOfVirtualIDDetailParam `json:"params"` // 详细参数
}

