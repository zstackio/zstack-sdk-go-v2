// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateVirtualRouter 更新VirtualRouter
func (cli *ZSClient) UpdateVirtualRouter(uuid string, params param.UpdateVirtualRouterParam) (*view.UpdateVirtualRouterEventView, error) {
	resp := view.UpdateVirtualRouterEventView{}
	if err := cli.Put("v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

