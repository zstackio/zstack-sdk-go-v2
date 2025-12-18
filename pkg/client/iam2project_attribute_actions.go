// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIAM2ProjectAttribute 查询IAM2ProjectAttribute列表
func (cli *ZSClient) QueryIAM2ProjectAttribute(params param.QueryParam) ([]view.QueryIAM2ProjectAttributeView, error) {
	var resp []view.QueryIAM2ProjectAttributeView
	return resp, cli.List("v1/iam2/projects/attributes", &params, &resp)
}

