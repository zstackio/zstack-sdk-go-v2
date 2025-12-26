// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetRouteTableVpcVRouterCandidate gets RouteTableVpcVRouterCandidate by uuid
func (cli *ZSClient) GetRouteTableVpcVRouterCandidate(uuid string) (*view.GetRouteTableVpcVRouterCandidateView, error) {
	var resp view.GetRouteTableVpcVRouterCandidateView
	if err := cli.Get("v1/vpc/virtual-routers/get-vpc-candidate", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
