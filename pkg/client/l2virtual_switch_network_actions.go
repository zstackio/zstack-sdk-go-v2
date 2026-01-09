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

func (cli *ZSClient) GetL2VirtualSwitchNetwork(uuid string) (*view.L2VirtualSwitchNetworkInventoryView, error) {
	var resp view.L2VirtualSwitchNetworkInventoryView
	if err := cli.Get("v1/l2-networks/virtual-switch", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
