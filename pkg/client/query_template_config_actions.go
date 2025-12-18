// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryTemplateConfig queries TemplateConfig list
func (cli *ZSClient) QueryTemplateConfig(params param.QueryParam) ([]view.TemplateConfigInventoryView, error) {
	var resp []view.TemplateConfigInventoryView
	return resp, cli.List("v1/template-configurations/configs", &params, &resp)
}
