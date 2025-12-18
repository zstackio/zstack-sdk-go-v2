// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySNSTextTemplate 查询SNSTextTemplate列表
func (cli *ZSClient) QuerySNSTextTemplate(params param.QueryParam) ([]view.QuerySNSTextTemplateView, error) {
	var resp []view.QuerySNSTextTemplateView
	return resp, cli.List("v1/zwatch/alarms/sns/text-templates", &params, &resp)
}

