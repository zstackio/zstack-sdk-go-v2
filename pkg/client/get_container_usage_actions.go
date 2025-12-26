// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetContainerUsage gets ContainerUsage by uuid
func (cli *ZSClient) GetContainerUsage(uuid string) (*view.GetContainerUsageView, error) {
	var resp view.GetContainerUsageView
	if err := cli.Get("v1/container/usage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
