// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateVirtualBorderRouterRemote updates VirtualBorderRouterRemote
func (cli *ZSClient) UpdateVirtualBorderRouterRemote(uuid string, params param.UpdateVirtualBorderRouterRemoteParam) (*view.UpdateVirtualBorderRouterRemoteEventView, error) {
	resp := view.UpdateVirtualBorderRouterRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/border-router/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
