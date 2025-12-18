// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryIAM2ProjectAccountRef 查询IAM2ProjectAccountRef列表
func (cli *ZSClient) QueryIAM2ProjectAccountRef(params param.QueryParam) ([]view.QueryIAM2ProjectAccountRefView, error) {
	var resp []view.QueryIAM2ProjectAccountRefView
	return resp, cli.List("v1/iam2/projects/account/refs", &params, &resp)
}

