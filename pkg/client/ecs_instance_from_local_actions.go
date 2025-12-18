// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEcsInstanceFromLocal 查询EcsInstanceFromLocal列表
func (cli *ZSClient) QueryEcsInstanceFromLocal(params param.QueryParam) ([]view.QueryEcsInstanceFromLocalView, error) {
	var resp []view.QueryEcsInstanceFromLocalView
	return resp, cli.List("v1/hybrid/aliyun/ecs", &params, &resp)
}

