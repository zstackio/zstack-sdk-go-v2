// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetManagementNodeDirCapacity gets ManagementNodeDirCapacity by uuid
func (cli *ZSClient) GetManagementNodeDirCapacity(uuid string) (*view.GetManagementNodeDirCapacityView, error) {
	var resp view.GetManagementNodeDirCapacityView
	if err := cli.Get("v1/zwatch/mn", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
