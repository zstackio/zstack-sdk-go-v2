// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAliyunRouteEntryFromLocal 查询AliyunRouteEntryFromLocal列表
func (cli *ZSClient) QueryAliyunRouteEntryFromLocal(params param.QueryParam) ([]view.QueryAliyunRouteEntryFromLocalView, error) {
	var resp []view.QueryAliyunRouteEntryFromLocalView
	return resp, cli.List("v1/hybrid/aliyun/route-entry", &params, &resp)
}

