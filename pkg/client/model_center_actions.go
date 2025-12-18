// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryModelCenter 查询ModelCenter列表
func (cli *ZSClient) QueryModelCenter(params param.QueryParam) ([]view.QueryModelCenterView, error) {
	var resp []view.QueryModelCenterView
	return resp, cli.List("v1/ai/model-centers", &params, &resp)
}

