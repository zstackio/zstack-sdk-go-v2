// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVirtualRouterSoftwareVersion gets VirtualRouterSoftwareVersion by uuid
func (cli *ZSClient) GetVirtualRouterSoftwareVersion(uuid string) (*view.GetVirtualRouterSoftwareVersionView, error) {
	var resp view.GetVirtualRouterSoftwareVersionView
	if err := cli.Get("v1/vpc/virtual-routers/softwareversion", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
