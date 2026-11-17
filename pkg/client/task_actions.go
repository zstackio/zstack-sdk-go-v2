// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// GetChainTask gets ChainTask by uuid
func (cli *ZSClient) GetChainTask() (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.GetWithRespKey("v1/core/task-details", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTaskProgress gets TaskProgress by uuid
func (cli *ZSClient) GetTaskProgress(uuid string) (*view.TaskProgressInventoryView, error) {
	var resp view.TaskProgressInventoryView
	if err := cli.GetWithRespKey("v1/task-progresses", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanQueue operates on Queue
func (cli *ZSClient) CleanQueue(params param.CleanQueueParam) (*view.CleanQueueEventView, error) {
	resp := view.CleanQueueEventView{}
	if err := cli.PutWithRespKey("v1/clean/queue", "", "", map[string]interface{}{
		"cleanQueue": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
