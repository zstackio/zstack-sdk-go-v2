// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryImageGroup 查询ImageGroup列表
func (cli *ZSClient) QueryImageGroup(params param.QueryParam) ([]view.QueryImageGroupView, error) {
	var resp []view.QueryImageGroupView
	return resp, cli.List("v1/imagegroups", &params, &resp)
}

