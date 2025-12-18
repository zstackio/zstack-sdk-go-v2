// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVirtualRouterOffering 创建VirtualRouterOffering
func (cli *ZSClient) CreateVirtualRouterOffering(params param.CreateVirtualRouterOfferingParam) (*view.CreateInstanceOfferingEventView, error) {
	resp := view.CreateInstanceOfferingEventView{}
	if err := cli.Post("v1/instance-offerings/virtual-routers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

