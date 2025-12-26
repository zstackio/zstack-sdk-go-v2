// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryManagementNode queries ManagementNode list
func (cli *ZSClient) QueryManagementNode(params *param.QueryParam) ([]view.ManagementNodeInventoryView, error) {
	var resp []view.ManagementNodeInventoryView
	return resp, cli.List("v1/management-nodes", params, &resp)
}
