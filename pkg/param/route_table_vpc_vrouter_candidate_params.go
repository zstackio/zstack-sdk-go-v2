// Copyright (c) ZStack.io, Inc.

package param

// GetRouteTableVpcVRouterCandidateDetailParam GetRouteTableVpcVRouterCandidate详细参数
type GetRouteTableVpcVRouterCandidateDetailParam struct {
	rest string `json:"tableUuid,omitempty"`
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetRouteTableVpcVRouterCandidateParam GetRouteTableVpcVRouterCandidate请求参数
type GetRouteTableVpcVRouterCandidateParam struct {
	BaseParam
	Params GetRouteTableVpcVRouterCandidateDetailParam `json:"params"` // 详细参数
}

