// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmRDP gets VmRDP by uuid
func (cli *ZSClient) GetVmRDP(uuid string) (*view.GetVmRDPView, error) {
	var resp view.GetVmRDPView
	if err := cli.Get("v1/vm-instances/{uuid}/rdp", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
