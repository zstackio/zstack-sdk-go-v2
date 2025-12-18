// Copyright (c) ZStack.io, Inc.

package param

// GetRouteTableVpcVRouterCandidateDetailParam GetRouteTableVpcVRouterCandidate detail param
type GetRouteTableVpcVRouterCandidateDetailParam struct {
	TableUuid string `json:"tableUuid,omitempty"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetRouteTableVpcVRouterCandidateParam GetRouteTableVpcVRouterCandidate request param
type GetRouteTableVpcVRouterCandidateParam struct {
	BaseParam
	Params GetRouteTableVpcVRouterCandidateDetailParam `json:"params"`
}
