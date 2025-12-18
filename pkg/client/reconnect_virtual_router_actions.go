// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ReconnectVirtualRouter operates on ReconnectVirtualRouter
func (cli *ZSClient) ReconnectVirtualRouter(uuid string, params param.ReconnectVirtualRouterParam) (*view.ReconnectVirtualRouterEventView, error) {
	resp := view.ReconnectVirtualRouterEventView{}
	if err := cli.Put("v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
