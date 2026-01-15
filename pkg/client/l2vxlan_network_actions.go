// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateL2VxlanNetwork creates L2VxlanNetwork
func (cli *ZSClient) CreateL2VxlanNetwork(params param.CreateL2VxlanNetworkParam) (*view.L2VxlanNetworkInventoryView, error) {
	resp := view.L2VxlanNetworkInventoryView{}
	if err := cli.Post("v1/l2-networks/vxlan", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryL2VxlanNetwork queries L2VxlanNetwork list
func (cli *ZSClient) QueryL2VxlanNetwork(params *param.QueryParam) ([]view.L2VxlanNetworkInventoryView, error) {
	var resp []view.L2VxlanNetworkInventoryView
	return resp, cli.List("v1/l2-networks/vxlan", params, &resp)
}

// PageL2VxlanNetwork Pagination
func (cli *ZSClient) PageL2VxlanNetwork(params *param.QueryParam) ([]view.L2VxlanNetworkInventoryView, int, error) {
	var l2VxlanNetworks []view.L2VxlanNetworkInventoryView
	total, err := cli.Page("v1/l2-networks/vxlan", params, &l2VxlanNetworks)
	return l2VxlanNetworks, total, err
}
