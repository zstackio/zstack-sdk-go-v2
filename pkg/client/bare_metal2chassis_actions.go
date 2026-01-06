// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// InspectBareMetal2Chassis operates on BareMetal2Chassis
func (cli *ZSClient) InspectBareMetal2Chassis(uuid string, params param.InspectBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.InspectBareMetal2ChassisEventView
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateBareMetal2Chassis updates BareMetal2Chassis
func (cli *ZSClient) UpdateBareMetal2Chassis(uuid string, params param.UpdateBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.UpdateBareMetal2ChassisEventView
	if err := cli.Put("v1/baremetal2/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryBareMetal2Chassis queries BareMetal2Chassis list
func (cli *ZSClient) QueryBareMetal2Chassis(params *param.QueryParam) ([]view.BareMetal2ChassisInventoryView, error) {
	var resp []view.BareMetal2ChassisInventoryView
	return resp, cli.List("v1/baremetal2/chassis", params, &resp)
}
// DeleteBareMetal2Chassis deletes BareMetal2Chassis
func (cli *ZSClient) DeleteBareMetal2Chassis(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/chassis/{uuid}", uuid, string(deleteMode))
}
