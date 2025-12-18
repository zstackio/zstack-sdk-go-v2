// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVpcVRouter creates VpcVRouter
func (cli *ZSClient) CreateVpcVRouter(params param.CreateVpcVRouterParam) (*view.CreateVpcVRouterEventView, error) {
	resp := view.CreateVpcVRouterEventView{}
	if err := cli.Post("v1/vpc/virtual-routers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
