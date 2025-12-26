// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddDnsToVpcRouter adds DnsToVpcRouter
func (cli *ZSClient) AddDnsToVpcRouter(params param.AddDnsToVpcRouterParam) (*view.AddDnsToVpcRouterEventView, error) {
	resp := view.AddDnsToVpcRouterEventView{}
	if err := cli.Post("v1/vpc/virtual-routers/{uuid}/dns", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
