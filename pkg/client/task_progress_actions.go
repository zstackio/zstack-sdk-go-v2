// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetTaskProgress 获取TaskProgress详情
func (cli *ZSClient) GetTaskProgress(uuid string) (*view.GetTaskProgressView, error) {
	var resp view.GetTaskProgressView
	if err := cli.Get("v1/task-progresses/{apiId}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

