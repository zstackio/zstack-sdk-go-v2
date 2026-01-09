// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// RebootBaremetalInstance operates on BaremetalInstance
func (cli *ZSClient) RebootBaremetalInstance(uuid string, params param.RebootBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	var resp view.RebootBaremetalInstanceEventView
	if err := cli.Put("v1/baremetal/instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// StartBaremetalInstance starts BaremetalInstance
func (cli *ZSClient) StartBaremetalInstance(uuid string, params param.StartBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	var resp view.StartBaremetalInstanceEventView
	if err := cli.Put("v1/baremetal/instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateBaremetalInstance creates BaremetalInstance
func (cli *ZSClient) CreateBaremetalInstance(params param.CreateBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	var resp view.CreateBaremetalInstanceEventView
	if err := cli.Post("v1/baremetal/instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DestroyBaremetalInstance destroys BaremetalInstance
func (cli *ZSClient) DestroyBaremetalInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal/instances", uuid, string(deleteMode))
}
// ExpungeBaremetalInstance operates on BaremetalInstance
func (cli *ZSClient) ExpungeBaremetalInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal/instances/actions", uuid, string(deleteMode))
}
// UpdateBaremetalInstance updates BaremetalInstance
func (cli *ZSClient) UpdateBaremetalInstance(uuid string, params param.UpdateBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	var resp view.UpdateBaremetalInstanceEventView
	if err := cli.Put("v1/baremetal/instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// StopBaremetalInstance stops BaremetalInstance
func (cli *ZSClient) StopBaremetalInstance(uuid string, params param.StopBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	var resp view.StopBaremetalInstanceEventView
	if err := cli.Put("v1/baremetal/instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryBaremetalInstance queries BaremetalInstance list
func (cli *ZSClient) QueryBaremetalInstance(params *param.QueryParam) ([]view.BaremetalInstanceInventoryView, error) {
	var resp []view.BaremetalInstanceInventoryView
	return resp, cli.List("v1/baremetal/instances", params, &resp)
}

func (cli *ZSClient) GetBaremetalInstance(uuid string) (*view.BaremetalInstanceInventoryView, error) {
	var resp view.BaremetalInstanceInventoryView
	if err := cli.Get("v1/baremetal/instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// RecoverBaremetalInstance operates on BaremetalInstance
func (cli *ZSClient) RecoverBaremetalInstance(uuid string, params param.RecoverBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	var resp view.RecoverBaremetalInstanceEventView
	if err := cli.Put("v1/baremetal/instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
