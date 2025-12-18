// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVCenterDatacenter 查询VCenterDatacenter列表
func (cli *ZSClient) QueryVCenterDatacenter(params param.QueryParam) ([]view.QueryVCenterDatacenterView, error) {
	var resp []view.QueryVCenterDatacenterView
	return resp, cli.List("v1/vcenters/datacenters", &params, &resp)
}

