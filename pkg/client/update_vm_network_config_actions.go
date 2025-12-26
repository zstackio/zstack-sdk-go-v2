// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVmNetworkConfig updates VmNetworkConfig
func (cli *ZSClient) UpdateVmNetworkConfig(uuid string, params param.UpdateVmNetworkConfigParam) (*view.UpdateVmNetworkConfigEventView, error) {
	resp := view.UpdateVmNetworkConfigEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/update-nic-config", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
