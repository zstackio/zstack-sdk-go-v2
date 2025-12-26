// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetHostMultipathTopology gets HostMultipathTopology by uuid
func (cli *ZSClient) GetHostMultipathTopology(uuid string) (*view.GetHostMultipathTopologyView, error) {
	var resp view.GetHostMultipathTopologyView
	if err := cli.Get("v1/storage-devices/multipath/topology", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
