// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryL2VirtualSwitchNetwork queries L2VirtualSwitchNetwork list
func (cli *ZSClient) QueryL2VirtualSwitchNetwork(params *param.QueryParam) ([]view.L2VirtualSwitchNetworkInventoryView, error) {
	var resp []view.L2VirtualSwitchNetworkInventoryView
	return resp, cli.List("v1/l2-networks/virtual-switch", params, &resp)
}

// PageL2VirtualSwitchNetwork Pagination
func (cli *ZSClient) PageL2VirtualSwitchNetwork(params *param.QueryParam) ([]view.L2VirtualSwitchNetworkInventoryView, int, error) {
	var l2VirtualSwitchNetworks []view.L2VirtualSwitchNetworkInventoryView
	total, err := cli.Page("v1/l2-networks/virtual-switch", params, &l2VirtualSwitchNetworks)
	return l2VirtualSwitchNetworks, total, err
}
