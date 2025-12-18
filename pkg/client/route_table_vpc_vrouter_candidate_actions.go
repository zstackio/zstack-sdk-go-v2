// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetRouteTableVpcVRouterCandidate 获取RouteTableVpcVRouterCandidate详情
func (cli *ZSClient) GetRouteTableVpcVRouterCandidate(uuid string) (*view.GetRouteTableVpcVRouterCandidateView, error) {
	var resp view.GetRouteTableVpcVRouterCandidateView
	if err := cli.Get("v1/vpc/virtual-routers/get-vpc-candidate", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

