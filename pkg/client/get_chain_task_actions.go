// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetChainTask gets ChainTask by uuid
func (cli *ZSClient) GetChainTask(uuid string) (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.Get("v1/core/task-details", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
