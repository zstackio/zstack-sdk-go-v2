// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryNativeCluster 查询NativeCluster列表
func (cli *ZSClient) QueryNativeCluster(params param.QueryParam) ([]view.QueryNativeClusterView, error) {
	var resp []view.QueryNativeClusterView
	return resp, cli.List("v1/container/native/cluster", &params, &resp)
}

