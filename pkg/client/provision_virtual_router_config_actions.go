// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ProvisionVirtualRouterConfig operates on ProvisionVirtualRouterConfig
func (cli *ZSClient) ProvisionVirtualRouterConfig(uuid string, params param.ProvisionVirtualRouterConfigParam) (*view.ProvisionVirtualRouterConfigEventView, error) {
	resp := view.ProvisionVirtualRouterConfigEventView{}
	if err := cli.Put("v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/provision", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
