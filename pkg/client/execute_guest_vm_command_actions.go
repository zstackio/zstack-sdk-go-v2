// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ExecuteGuestVmCommand operates on ExecuteGuestVmCommand
func (cli *ZSClient) ExecuteGuestVmCommand(params param.ExecuteGuestVmCommandParam) (*view.ExecuteGuestVmCommandEventView, error) {
	resp := view.ExecuteGuestVmCommandEventView{}
	if err := cli.Post("v1/vm-instances/commands/exec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
