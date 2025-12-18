// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ProvisionVirtualRouterConfig 操作ProvisionVirtualRouterConfig
func (cli *ZSClient) ProvisionVirtualRouterConfig(uuid string, params param.ProvisionVirtualRouterConfigParam) (*view.ProvisionVirtualRouterConfigEventView, error) {
	resp := view.ProvisionVirtualRouterConfigEventView{}
	if err := cli.Put("v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/provision", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

