// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcSnatState 查询VpcSnatState列表
func (cli *ZSClient) QueryVpcSnatState(params param.QueryParam) ([]view.QueryVpcSnatStateView, error) {
	var resp []view.QueryVpcSnatStateView
	return resp, cli.List("v1/vpc/virtual-routers/networkservicestate/snat", &params, &resp)
}

