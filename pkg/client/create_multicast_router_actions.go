// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateMulticastRouter creates MulticastRouter
func (cli *ZSClient) CreateMulticastRouter(params param.CreateMulticastRouterParam) (*view.CreateMulticastRouterEventView, error) {
	resp := view.CreateMulticastRouterEventView{}
	if err := cli.Post("v1/multicast/virtual-routers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
