// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryZBox 查询ZBox列表
func (cli *ZSClient) QueryZBox(params param.QueryParam) ([]view.QueryZBoxView, error) {
	var resp []view.QueryZBoxView
	return resp, cli.List("v1/zbox", &params, &resp)
}

