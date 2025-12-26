// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmMonitorNumber operates on SetVmMonitorNumber
func (cli *ZSClient) SetVmMonitorNumber(uuid string, params param.SetVmMonitorNumberParam) (*view.SetVmMonitorNumberEventView, error) {
	resp := view.SetVmMonitorNumberEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
