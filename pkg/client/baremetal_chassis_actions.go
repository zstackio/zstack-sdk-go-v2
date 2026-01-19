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

// PageBaremetalChassis Pagination
func (cli *ZSClient) PageBaremetalChassis(params *param.QueryParam) ([]view.BaremetalChassisInventoryView, int, error) {
	var baremetalChassis []view.BaremetalChassisInventoryView
	total, err := cli.Page("v1/baremetal/chassis", params, &baremetalChassis)
	return baremetalChassis, total, err
}
// InspectBaremetalChassis operates on BaremetalChassis
func (cli *ZSClient) InspectBaremetalChassis(uuid string, params param.InspectBaremetalChassisParam) (*view.BaremetalChassisInventoryView, error) {
	resp := view.BaremetalChassisInventoryView{}
	if err := cli.Put("v1/baremetal/chassis", uuid, map[string]interface{}{
		"inspectBaremetalChassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateBaremetalChassis updates BaremetalChassis
func (cli *ZSClient) UpdateBaremetalChassis(uuid string, params param.UpdateBaremetalChassisParam) (*view.BaremetalChassisInventoryView, error) {
	resp := view.BaremetalChassisInventoryView{}
	if err := cli.Put("v1/baremetal/chassis", uuid, map[string]interface{}{
		"updateBaremetalChassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteBaremetalChassis deletes BaremetalChassis
func (cli *ZSClient) DeleteBaremetalChassis(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal/chassis", uuid, string(deleteMode))
}
// CreateBaremetalChassis creates BaremetalChassis
func (cli *ZSClient) CreateBaremetalChassis(params param.CreateBaremetalChassisParam) (*view.BaremetalChassisInventoryView, error) {
	resp := view.BaremetalChassisInventoryView{}
	if err := cli.Post("v1/baremetal/chassis", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
