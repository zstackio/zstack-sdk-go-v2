// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmConsolePassword gets VmConsolePassword by uuid
func (cli *ZSClient) GetVmConsolePassword(uuid string) (*view.GetVmConsolePasswordView, error) {
	var resp view.GetVmConsolePasswordView
	if err := cli.Get("v1/vm-instances/{uuid}/console-passwords", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
