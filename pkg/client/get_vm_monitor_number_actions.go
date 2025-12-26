// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmMonitorNumber gets VmMonitorNumber by uuid
func (cli *ZSClient) GetVmMonitorNumber(uuid string) (*view.GetVmMonitorNumberView, error) {
	var resp view.GetVmMonitorNumberView
	if err := cli.Get("v1/vm-instances/{uuid}/monitorNumber", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
