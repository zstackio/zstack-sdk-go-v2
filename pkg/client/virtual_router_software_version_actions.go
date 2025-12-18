// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVirtualRouterSoftwareVersion 获取VirtualRouterSoftwareVersion详情
func (cli *ZSClient) GetVirtualRouterSoftwareVersion(uuid string) (*view.GetVirtualRouterSoftwareVersionView, error) {
	var resp view.GetVirtualRouterSoftwareVersionView
	if err := cli.Get("v1/vpc/virtual-routers/softwareversion", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

