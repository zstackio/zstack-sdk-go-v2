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
	resp := view.BaremetalInstanceInventoryView{}
	if err := cli.Put("v1/baremetal/instances", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StartBaremetalInstance starts BaremetalInstance
func (cli *ZSClient) StartBaremetalInstance(uuid string, params param.StartBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	resp := view.BaremetalInstanceInventoryView{}
	if err := cli.Put("v1/baremetal/instances", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateBaremetalInstance creates BaremetalInstance
func (cli *ZSClient) CreateBaremetalInstance(params param.CreateBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	resp := view.BaremetalInstanceInventoryView{}
	if err := cli.Post("v1/baremetal/instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DestroyBaremetalInstance destroys BaremetalInstance
func (cli *ZSClient) DestroyBaremetalInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal/instances", uuid, string(deleteMode))
}
// ExpungeBaremetalInstance operates on BaremetalInstance
func (cli *ZSClient) ExpungeBaremetalInstance(uuid string) error {
	params := map[string]interface{}{
		"expungeBaremetalInstance": map[string]interface{}{},
	}
	return cli.Put("v1/baremetal/instances", uuid, params, nil)
}
// UpdateBaremetalInstance updates BaremetalInstance
func (cli *ZSClient) UpdateBaremetalInstance(uuid string, params param.UpdateBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	resp := view.BaremetalInstanceInventoryView{}
	if err := cli.Put("v1/baremetal/instances", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StopBaremetalInstance stops BaremetalInstance
func (cli *ZSClient) StopBaremetalInstance(uuid string, params param.StopBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	resp := view.BaremetalInstanceInventoryView{}
	if err := cli.Put("v1/baremetal/instances", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryBaremetalInstance queries BaremetalInstance list
func (cli *ZSClient) QueryBaremetalInstance(params *param.QueryParam) ([]view.BaremetalInstanceInventoryView, error) {
	var resp []view.BaremetalInstanceInventoryView
	return resp, cli.List("v1/baremetal/instances", params, &resp)
}

// PageBaremetalInstance Pagination
func (cli *ZSClient) PageBaremetalInstance(params *param.QueryParam) ([]view.BaremetalInstanceInventoryView, int, error) {
	var baremetalInstances []view.BaremetalInstanceInventoryView
	total, err := cli.Page("v1/baremetal/instances", params, &baremetalInstances)
	return baremetalInstances, total, err
}
// RecoverBaremetalInstance operates on BaremetalInstance
func (cli *ZSClient) RecoverBaremetalInstance(uuid string, params param.RecoverBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	resp := view.BaremetalInstanceInventoryView{}
	if err := cli.Put("v1/baremetal/instances", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
