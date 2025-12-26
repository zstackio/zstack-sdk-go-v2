// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmvNUMATopology gets VmvNUMATopology by uuid
func (cli *ZSClient) GetVmvNUMATopology(uuid string) (*view.GetVmvNUMATopologyView, error) {
	var resp view.GetVmvNUMATopologyView
	if err := cli.Get("v1/vm-instances/{uuid}/vnuma-topology", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
