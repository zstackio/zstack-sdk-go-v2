// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryModel 查询Model列表
func (cli *ZSClient) QueryModel(params param.QueryParam) ([]view.QueryModelView, error) {
	var resp []view.QueryModelView
	return resp, cli.List("v1/ai/models", &params, &resp)
}

