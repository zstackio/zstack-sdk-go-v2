// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetHostNUMATopology gets HostNUMATopology by uuid
func (cli *ZSClient) GetHostNUMATopology(uuid string) (*view.GetHostNUMATopologyEventView, error) {
	var resp view.GetHostNUMATopologyEventView
	if err := cli.Get("v1/hosts/{uuid}/numa", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
