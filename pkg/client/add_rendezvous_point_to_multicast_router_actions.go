// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddRendezvousPointToMulticastRouter adds RendezvousPointToMulticastRouter
func (cli *ZSClient) AddRendezvousPointToMulticastRouter(params param.AddRendezvousPointToMulticastRouterParam) (*view.AddRendezvousPointToMulticastRouterEventView, error) {
	resp := view.AddRendezvousPointToMulticastRouterEventView{}
	if err := cli.Post("v1/multicast/virtual-routers/{uuid}/RendezvousPoint", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
