// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReconnectVirtualRouter operates on ReconnectVirtualRouter
func (cli *ZSClient) ReconnectVirtualRouter(uuid string, params param.ReconnectVirtualRouterParam) (*view.ReconnectVirtualRouterEventView, error) {
	resp := view.ReconnectVirtualRouterEventView{}
	if err := cli.Put("v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
