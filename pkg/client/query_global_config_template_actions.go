// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryGlobalConfigTemplate queries GlobalConfigTemplate list
func (cli *ZSClient) QueryGlobalConfigTemplate(params param.QueryParam) ([]view.GlobalConfigTemplateInventoryView, error) {
	var resp []view.GlobalConfigTemplateInventoryView
	return resp, cli.List("v1/template-configurations/templates", &params, &resp)
}
