// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVirtualRouter updates VirtualRouter
func (cli *ZSClient) UpdateVirtualRouter(uuid string, params param.UpdateVirtualRouterParam) (*view.UpdateVirtualRouterEventView, error) {
	resp := view.UpdateVirtualRouterEventView{}
	if err := cli.Put("v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
