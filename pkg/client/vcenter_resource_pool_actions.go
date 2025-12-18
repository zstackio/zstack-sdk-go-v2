// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVCenterResourcePool 查询VCenterResourcePool列表
func (cli *ZSClient) QueryVCenterResourcePool(params param.QueryParam) ([]view.QueryVCenterResourcePoolView, error) {
	var resp []view.QueryVCenterResourcePoolView
	return resp, cli.List("v1/vcenters/clusters/resourcepools", &params, &resp)
}

