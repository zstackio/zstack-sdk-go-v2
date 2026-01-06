// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddXDragonHost adds XDragonHost
func (cli *ZSClient) AddXDragonHost(params param.AddXDragonHostParam) (*view.HostInventoryView, error) {
	var resp view.AddHostEventView
	if err := cli.Post("v1/hosts/xdragon", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
