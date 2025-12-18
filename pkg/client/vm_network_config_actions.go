// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateVmNetworkConfig 更新VmNetworkConfig
func (cli *ZSClient) UpdateVmNetworkConfig(uuid string, params param.UpdateVmNetworkConfigParam) (*view.UpdateVmNetworkConfigEventView, error) {
	resp := view.UpdateVmNetworkConfigEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/update-nic-config", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

