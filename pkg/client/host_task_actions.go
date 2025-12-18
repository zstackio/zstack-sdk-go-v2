// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetHostTask 获取HostTask详情
func (cli *ZSClient) GetHostTask(uuid string) (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.Get("v1/hosts/task-details", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

