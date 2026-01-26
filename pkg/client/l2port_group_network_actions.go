// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryL2PortGroupNetwork queries L2PortGroupNetwork list
func (cli *ZSClient) QueryL2PortGroupNetwork(params *param.QueryParam) ([]view.L2PortGroupNetworkInventoryView, error) {
	var resp []view.L2PortGroupNetworkInventoryView
	return resp, cli.List("v1/l2-networks/port-group", params, &resp)
}

func (cli *ZSClient) GetL2PortGroupNetwork(uuid string) (*view.L2PortGroupNetworkInventoryView, error) {
	var resp view.L2PortGroupNetworkInventoryView
	if err := cli.Get("v1/l2-networks/port-group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageL2PortGroupNetwork Pagination
func (cli *ZSClient) PageL2PortGroupNetwork(params *param.QueryParam) ([]view.L2PortGroupNetworkInventoryView, int, error) {
	var l2PortGroupNetworks []view.L2PortGroupNetworkInventoryView
	total, err := cli.Page("v1/l2-networks/port-group", params, &l2PortGroupNetworks)
	return l2PortGroupNetworks, total, err
}
