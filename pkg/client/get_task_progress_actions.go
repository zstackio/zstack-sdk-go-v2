// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetTaskProgress gets TaskProgress by uuid
func (cli *ZSClient) GetTaskProgress(uuid string) (*view.GetTaskProgressView, error) {
	var resp view.GetTaskProgressView
	if err := cli.Get("v1/task-progresses/{apiId}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
