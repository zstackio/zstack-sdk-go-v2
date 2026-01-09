// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateL2VlanNetwork creates L2VlanNetwork
func (cli *ZSClient) CreateL2VlanNetwork(params param.CreateL2VlanNetworkParam) (*view.L2VlanNetworkInventoryView, error) {
	resp := view.L2VlanNetworkInventoryView{}
	if err := cli.Post("v1/l2-networks/vlan", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryL2VlanNetwork queries L2VlanNetwork list
func (cli *ZSClient) QueryL2VlanNetwork(params *param.QueryParam) ([]view.L2VlanNetworkInventoryView, error) {
	var resp []view.L2VlanNetworkInventoryView
	return resp, cli.List("v1/l2-networks/vlan", params, &resp)
}

func (cli *ZSClient) GetL2VlanNetwork(uuid string) (*view.L2VlanNetworkInventoryView, error) {
	var resp view.L2VlanNetworkInventoryView
	if err := cli.Get("v1/l2-networks/vlan", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
