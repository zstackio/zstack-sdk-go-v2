// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryConnectionAccessPointFromLocal 查询ConnectionAccessPointFromLocal列表
func (cli *ZSClient) QueryConnectionAccessPointFromLocal(params param.QueryParam) ([]view.QueryConnectionAccessPointFromLocalView, error) {
	var resp []view.QueryConnectionAccessPointFromLocalView
	return resp, cli.List("v1/hybrid/aliyun/access-point", &params, &resp)
}

