// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVirtualRouterSoftwareVersion gets VirtualRouterSoftwareVersion by uuid
func (cli *ZSClient) GetVirtualRouterSoftwareVersion(uuid string) (*view.GetVirtualRouterSoftwareVersionView, error) {
	var resp view.GetVirtualRouterSoftwareVersionView
	if err := cli.Get("v1/vpc/virtual-routers/softwareversion", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
