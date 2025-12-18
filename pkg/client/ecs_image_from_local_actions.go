// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEcsImageFromLocal 查询EcsImageFromLocal列表
func (cli *ZSClient) QueryEcsImageFromLocal(params param.QueryParam) ([]view.QueryEcsImageFromLocalView, error) {
	var resp []view.QueryEcsImageFromLocalView
	return resp, cli.List("v1/hybrid/aliyun/image", &params, &resp)
}

