// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAliyunVirtualRouterFromLocal 查询AliyunVirtualRouterFromLocal列表
func (cli *ZSClient) QueryAliyunVirtualRouterFromLocal(params param.QueryParam) ([]view.QueryAliyunVirtualRouterFromLocalView, error) {
	var resp []view.QueryAliyunVirtualRouterFromLocalView
	return resp, cli.List("v1/hybrid/aliyun/vrouter", &params, &resp)
}

