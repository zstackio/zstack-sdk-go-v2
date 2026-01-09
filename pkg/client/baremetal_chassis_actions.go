// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBaremetalChassis queries BaremetalChassis list
func (cli *ZSClient) QueryBaremetalChassis(params *param.QueryParam) ([]view.BaremetalChassisInventoryView, error) {
	var resp []view.BaremetalChassisInventoryView
	return resp, cli.List("v1/baremetal/chassis", params, &resp)
}

func (cli *ZSClient) GetBaremetalChassis(uuid string) (*view.BaremetalChassisInventoryView, error) {
	var resp view.BaremetalChassisInventoryView
	if err := cli.Get("v1/baremetal/chassis", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// InspectBaremetalChassis operates on BaremetalChassis
func (cli *ZSClient) InspectBaremetalChassis(uuid string, params param.InspectBaremetalChassisParam) (*view.BaremetalChassisInventoryView, error) {
	var resp view.InspectBaremetalChassisEventView
	if err := cli.Put("v1/baremetal/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateBaremetalChassis updates BaremetalChassis
func (cli *ZSClient) UpdateBaremetalChassis(uuid string, params param.UpdateBaremetalChassisParam) (*view.BaremetalChassisInventoryView, error) {
	var resp view.UpdateBaremetalChassisEventView
	if err := cli.Put("v1/baremetal/chassis/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteBaremetalChassis deletes BaremetalChassis
func (cli *ZSClient) DeleteBaremetalChassis(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal/chassis", uuid, string(deleteMode))
}
// CreateBaremetalChassis creates BaremetalChassis
func (cli *ZSClient) CreateBaremetalChassis(params param.CreateBaremetalChassisParam) (*view.BaremetalChassisInventoryView, error) {
	var resp view.CreateBaremetalChassisEventView
	if err := cli.Post("v1/baremetal/chassis", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
