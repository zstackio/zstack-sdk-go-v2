// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVpcSnatState queries VpcSnatState list
func (cli *ZSClient) QueryVpcSnatState(params *param.QueryParam) ([]view.VpcSnatStateInventoryView, error) {
	var resp []view.VpcSnatStateInventoryView
	return resp, cli.List("v1/vpc/virtual-routers/networkservicestate/snat", params, &resp)
}
