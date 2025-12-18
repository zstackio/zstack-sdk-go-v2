// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAliyunRouterInterfaceFromLocal 查询AliyunRouterInterfaceFromLocal列表
func (cli *ZSClient) QueryAliyunRouterInterfaceFromLocal(params param.QueryParam) ([]view.QueryAliyunRouterInterfaceFromLocalView, error) {
	var resp []view.QueryAliyunRouterInterfaceFromLocalView
	return resp, cli.List("v1/hybrid/aliyun/router-interface", &params, &resp)
}

