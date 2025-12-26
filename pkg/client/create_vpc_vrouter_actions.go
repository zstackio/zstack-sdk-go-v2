// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVpcVRouter creates VpcVRouter
func (cli *ZSClient) CreateVpcVRouter(params param.CreateVpcVRouterParam) (*view.CreateVpcVRouterEventView, error) {
	resp := view.CreateVpcVRouterEventView{}
	if err := cli.Post("v1/vpc/virtual-routers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
