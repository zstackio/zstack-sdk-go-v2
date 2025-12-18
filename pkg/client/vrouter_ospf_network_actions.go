// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVRouterOspfNetwork 查询VRouterOspfNetwork列表
func (cli *ZSClient) QueryVRouterOspfNetwork(params param.QueryParam) ([]view.QueryVRouterOspfNetworkView, error) {
	var resp []view.QueryVRouterOspfNetworkView
	return resp, cli.List("v1/routerArea/network", &params, &resp)
}

