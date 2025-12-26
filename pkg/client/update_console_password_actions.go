// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateConsolePassword updates ConsolePassword
func (cli *ZSClient) UpdateConsolePassword(uuid string, params param.UpdateConsolePasswordParam) (*view.UpdateConsolePasswordEventView, error) {
	resp := view.UpdateConsolePasswordEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
