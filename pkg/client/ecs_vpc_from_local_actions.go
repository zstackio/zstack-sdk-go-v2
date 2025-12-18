// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEcsVpcFromLocal 查询EcsVpcFromLocal列表
func (cli *ZSClient) QueryEcsVpcFromLocal(params param.QueryParam) ([]view.QueryEcsVpcFromLocalView, error) {
	var resp []view.QueryEcsVpcFromLocalView
	return resp, cli.List("v1/hybrid/aliyun/vpc", &params, &resp)
}

