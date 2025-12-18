// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetManagementNodeOS 获取ManagementNodeOS详情
func (cli *ZSClient) GetManagementNodeOS(uuid string) (*view.GetManagementNodeOSView, error) {
	var resp view.GetManagementNodeOSView
	if err := cli.Get("v1/management/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

