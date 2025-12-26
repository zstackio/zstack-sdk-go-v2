// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVRouterOspfArea queries VRouterOspfArea list
func (cli *ZSClient) QueryVRouterOspfArea(params *param.QueryParam) ([]view.RouterAreaInventoryView, error) {
	var resp []view.RouterAreaInventoryView
	return resp, cli.List("v1/routerArea", params, &resp)
}
