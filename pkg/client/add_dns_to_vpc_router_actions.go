// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddDnsToVpcRouter 操作AddDnsToVpcRouter
func (cli *ZSClient) AddDnsToVpcRouter(params param.AddDnsToVpcRouterParam) (*view.AddDnsToVpcRouterEventView, error) {
	resp := view.AddDnsToVpcRouterEventView{}
	if err := cli.Post("v1/vpc/virtual-routers/{uuid}/dns", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

