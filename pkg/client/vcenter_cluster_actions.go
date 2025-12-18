// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVCenterCluster 查询VCenterCluster列表
func (cli *ZSClient) QueryVCenterCluster(params param.QueryParam) ([]view.QueryVCenterClusterView, error) {
	var resp []view.QueryVCenterClusterView
	return resp, cli.List("v1/vcenters/clusters", &params, &resp)
}

