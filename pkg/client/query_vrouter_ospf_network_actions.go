// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVRouterOspfNetwork queries VRouterOspfNetwork list
func (cli *ZSClient) QueryVRouterOspfNetwork(params *param.QueryParam) ([]view.NetworkRouterAreaRefInventoryView, error) {
	var resp []view.NetworkRouterAreaRefInventoryView
	return resp, cli.List("v1/routerArea/network", params, &resp)
}
