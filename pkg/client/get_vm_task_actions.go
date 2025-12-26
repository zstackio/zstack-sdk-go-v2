// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVmTask gets VmTask by uuid
func (cli *ZSClient) GetVmTask(uuid string) (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.Get("v1/vm-instances/task-details", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
