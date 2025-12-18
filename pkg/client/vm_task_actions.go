// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVmTask 获取VmTask详情
func (cli *ZSClient) GetVmTask(uuid string) (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.Get("v1/vm-instances/task-details", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

