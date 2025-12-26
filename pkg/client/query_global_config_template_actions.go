// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryGlobalConfigTemplate queries GlobalConfigTemplate list
func (cli *ZSClient) QueryGlobalConfigTemplate(params *param.QueryParam) ([]view.GlobalConfigTemplateInventoryView, error) {
	var resp []view.GlobalConfigTemplateInventoryView
	return resp, cli.List("v1/template-configurations/templates", params, &resp)
}
