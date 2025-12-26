// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetHostResourceAllocation gets HostResourceAllocation by uuid
func (cli *ZSClient) GetHostResourceAllocation(uuid string) (*view.GetHostResourceAllocationEventView, error) {
	var resp view.GetHostResourceAllocationEventView
	if err := cli.Get("v1/hosts/{uuid}/resource-allocation", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
