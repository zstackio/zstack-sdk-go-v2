// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryTemplateConfig queries TemplateConfig list
func (cli *ZSClient) QueryTemplateConfig(params *param.QueryParam) ([]view.TemplateConfigInventoryView, error) {
	var resp []view.TemplateConfigInventoryView
	return resp, cli.List("v1/template-configurations/configs", params, &resp)
}
