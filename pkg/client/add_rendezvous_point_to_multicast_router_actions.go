// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddRendezvousPointToMulticastRouter adds RendezvousPointToMulticastRouter
func (cli *ZSClient) AddRendezvousPointToMulticastRouter(params param.AddRendezvousPointToMulticastRouterParam) (*view.AddRendezvousPointToMulticastRouterEventView, error) {
	resp := view.AddRendezvousPointToMulticastRouterEventView{}
	if err := cli.Post("v1/multicast/virtual-routers/{uuid}/RendezvousPoint", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
