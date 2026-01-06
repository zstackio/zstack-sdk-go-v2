// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryL2PortGroupNetwork queries L2PortGroupNetwork list
func (cli *ZSClient) QueryL2PortGroupNetwork(params *param.QueryParam) ([]view.L2PortGroupNetworkInventoryView, error) {
	var resp []view.L2PortGroupNetworkInventoryView
	return resp, cli.List("v1/l2-networks/port-group", params, &resp)
}
