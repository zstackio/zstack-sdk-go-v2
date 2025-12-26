// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryStackTemplate queries StackTemplate list
func (cli *ZSClient) QueryStackTemplate(params *param.QueryParam) ([]view.StackTemplateInventoryView, error) {
	var resp []view.StackTemplateInventoryView
	return resp, cli.List("v1/cloudformation/template", params, &resp)
}
