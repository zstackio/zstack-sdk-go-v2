// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetManagementNodeDirCapacity 获取ManagementNodeDirCapacity详情
func (cli *ZSClient) GetManagementNodeDirCapacity(uuid string) (*view.GetManagementNodeDirCapacityView, error) {
	var resp view.GetManagementNodeDirCapacityView
	if err := cli.Get("v1/zwatch/mn", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

