// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmConsolePassword operates on SetVmConsolePassword
func (cli *ZSClient) SetVmConsolePassword(uuid string, params param.SetVmConsolePasswordParam) (*view.SetVmConsolePasswordEventView, error) {
	resp := view.SetVmConsolePasswordEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
