// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryManagementNode 查询ManagementNode列表
func (cli *ZSClient) QueryManagementNode(params param.QueryParam) ([]view.QueryManagementNodeView, error) {
	var resp []view.QueryManagementNodeView
	return resp, cli.List("v1/management-nodes", &params, &resp)
}

