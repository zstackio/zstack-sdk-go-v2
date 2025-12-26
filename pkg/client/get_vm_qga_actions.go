// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmQga gets VmQga by uuid
func (cli *ZSClient) GetVmQga(uuid string) (*view.GetVmQgaView, error) {
	var resp view.GetVmQgaView
	if err := cli.Get("v1/vm-instances/{uuid}/qga", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
