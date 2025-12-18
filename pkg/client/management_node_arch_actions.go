// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetManagementNodeArch 获取ManagementNodeArch详情
func (cli *ZSClient) GetManagementNodeArch(uuid string) (*view.GetManagementNodeArchView, error) {
	var resp view.GetManagementNodeArchView
	if err := cli.Get("v1/management-nodes/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

