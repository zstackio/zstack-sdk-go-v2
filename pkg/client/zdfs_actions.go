// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryZdfs 查询Zdfs列表
func (cli *ZSClient) QueryZdfs(params param.QueryParam) ([]view.QueryZdfsView, error) {
	var resp []view.QueryZdfsView
	return resp, cli.List("v1/zdfs", &params, &resp)
}

