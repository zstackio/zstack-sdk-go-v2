// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVtep 查询Vtep列表
func (cli *ZSClient) QueryVtep(params param.QueryParam) ([]view.QueryVtepView, error) {
	var resp []view.QueryVtepView
	return resp, cli.List("v1/l2-networks/vteps", &params, &resp)
}

