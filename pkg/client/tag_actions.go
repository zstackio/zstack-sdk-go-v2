// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryTag 查询Tag列表
func (cli *ZSClient) QueryTag(params param.QueryParam) ([]view.QueryTagView, error) {
	var resp []view.QueryTagView
	return resp, cli.List("v1/tags", &params, &resp)
}

