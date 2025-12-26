// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmConsoleMode operates on SetVmConsoleMode
func (cli *ZSClient) SetVmConsoleMode(uuid string, params param.SetVmConsoleModeParam) (*view.SetVmConsoleModeEventView, error) {
	resp := view.SetVmConsoleModeEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
