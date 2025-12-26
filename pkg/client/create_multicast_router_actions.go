// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateMulticastRouter creates MulticastRouter
func (cli *ZSClient) CreateMulticastRouter(params param.CreateMulticastRouterParam) (*view.CreateMulticastRouterEventView, error) {
	resp := view.CreateMulticastRouterEventView{}
	if err := cli.Post("v1/multicast/virtual-routers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
