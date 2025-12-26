// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVirtualRouterOffering creates VirtualRouterOffering
func (cli *ZSClient) CreateVirtualRouterOffering(params param.CreateVirtualRouterOfferingParam) (*view.CreateInstanceOfferingEventView, error) {
	resp := view.CreateInstanceOfferingEventView{}
	if err := cli.Post("v1/instance-offerings/virtual-routers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
