// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateL2VxlanNetworkPool creates L2VxlanNetworkPool
func (cli *ZSClient) CreateL2VxlanNetworkPool(params param.CreateL2VxlanNetworkPoolParam) (*view.L2VxlanNetworkPoolInventoryView, error) {
	resp := view.L2VxlanNetworkPoolInventoryView{}
	if err := cli.Post("v1/l2-networks/vxlan-pool", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryL2VxlanNetworkPool queries L2VxlanNetworkPool list
func (cli *ZSClient) QueryL2VxlanNetworkPool(params *param.QueryParam) ([]view.L2VxlanNetworkPoolInventoryView, error) {
	var resp []view.L2VxlanNetworkPoolInventoryView
	return resp, cli.List("v1/l2-networks/vxlan-pool", params, &resp)
}
