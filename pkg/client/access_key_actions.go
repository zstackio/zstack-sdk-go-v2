// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAccessKey 查询AccessKey列表
func (cli *ZSClient) QueryAccessKey(params param.QueryParam) ([]view.QueryAccessKeyView, error) {
	var resp []view.QueryAccessKeyView
	return resp, cli.List("v1/accesskeys", &params, &resp)
}

