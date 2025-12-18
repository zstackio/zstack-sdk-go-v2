// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIAM2ProjectTemplate 查询IAM2ProjectTemplate列表
func (cli *ZSClient) QueryIAM2ProjectTemplate(params param.QueryParam) ([]view.QueryIAM2ProjectTemplateView, error) {
	var resp []view.QueryIAM2ProjectTemplateView
	return resp, cli.List("v1/iam2/projects/templates", &params, &resp)
}

