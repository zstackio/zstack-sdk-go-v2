// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateVirtualRouterSoftwareVersion updates VirtualRouterSoftwareVersion
func (cli *ZSClient) UpdateVirtualRouterSoftwareVersion(uuid string, params param.UpdateVirtualRouterSoftwareVersionParam) (*view.UpdateVirtualRouterSoftwareVersionEventView, error) {
	resp := view.UpdateVirtualRouterSoftwareVersionEventView{}
	if err := cli.Put("v1/vpc/virtual-routers/{uuid}/softwareversion", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
