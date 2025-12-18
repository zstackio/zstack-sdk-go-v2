// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryGlobalConfigTemplate 查询GlobalConfigTemplate列表
func (cli *ZSClient) QueryGlobalConfigTemplate(params param.QueryParam) ([]view.QueryGlobalConfigTemplateView, error) {
	var resp []view.QueryGlobalConfigTemplateView
	return resp, cli.List("v1/template-configurations/templates", &params, &resp)
}

