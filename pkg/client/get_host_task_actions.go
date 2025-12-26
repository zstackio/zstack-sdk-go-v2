// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetHostTask gets HostTask by uuid
func (cli *ZSClient) GetHostTask(uuid string) (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.Get("v1/hosts/task-details", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
