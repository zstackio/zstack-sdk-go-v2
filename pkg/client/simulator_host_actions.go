// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddSimulatorHost adds SimulatorHost
func (cli *ZSClient) AddSimulatorHost(params param.AddSimulatorHostParam) (*view.HostInventoryView, error) {
	var resp view.AddHostEventView
	if err := cli.Post("v1/hosts/simulators", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
