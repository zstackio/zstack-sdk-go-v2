// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmConsoleAddress gets VmConsoleAddress by uuid
func (cli *ZSClient) GetVmConsoleAddress(uuid string) (*view.GetVmConsoleAddressView, error) {
	var resp view.GetVmConsoleAddressView
	if err := cli.Get("v1/vm-instances/{uuid}/console-addresses", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
