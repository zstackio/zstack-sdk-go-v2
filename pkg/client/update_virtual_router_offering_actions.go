// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateVirtualRouterOffering updates VirtualRouterOffering
func (cli *ZSClient) UpdateVirtualRouterOffering(uuid string, params param.UpdateVirtualRouterOfferingParam) (*view.UpdateInstanceOfferingEventView, error) {
	resp := view.UpdateInstanceOfferingEventView{}
	if err := cli.Put("v1/instance-offerings/virtual-routers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
