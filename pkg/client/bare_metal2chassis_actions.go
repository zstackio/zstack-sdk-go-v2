// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// InspectBareMetal2Chassis operates on BareMetal2Chassis
func (cli *ZSClient) InspectBareMetal2Chassis(uuid string, params param.InspectBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/chassis", uuid, "", map[string]interface{}{
		"inspectBareMetal2Chassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateBareMetal2Chassis updates BareMetal2Chassis
func (cli *ZSClient) UpdateBareMetal2Chassis(uuid string, params param.UpdateBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/chassis", uuid, "", map[string]interface{}{
		"updateBareMetal2Chassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryBareMetal2Chassis queries BareMetal2Chassis list
func (cli *ZSClient) QueryBareMetal2Chassis(params *param.QueryParam) ([]view.BareMetal2ChassisInventoryView, error) {
	var resp []view.BareMetal2ChassisInventoryView
	return resp, cli.List("v1/baremetal2/chassis", params, &resp)
}

func (cli *ZSClient) GetBareMetal2Chassis(uuid string) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.BareMetal2ChassisInventoryView
	if err := cli.Get("v1/baremetal2/chassis", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBareMetal2Chassis Pagination
func (cli *ZSClient) PageBareMetal2Chassis(params *param.QueryParam) ([]view.BareMetal2ChassisInventoryView, int, error) {
	var bareMetal2Chassis []view.BareMetal2ChassisInventoryView
	total, err := cli.Page("v1/baremetal2/chassis", params, &bareMetal2Chassis)
	return bareMetal2Chassis, total, err
}
// DeleteBareMetal2Chassis deletes BareMetal2Chassis
func (cli *ZSClient) DeleteBareMetal2Chassis(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/chassis", uuid, string(deleteMode))
}

// PowerOffBareMetal2Chassis operates on PowerOffBareMetal2Chassis
func (cli *ZSClient) PowerOffBareMetal2Chassis(uuid string, params param.PowerOffBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/chassis", uuid, "", map[string]interface{}{
		"powerOffBareMetal2Chassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBareMetal2ChassisPowerStatus gets BareMetal2ChassisPowerStatus by uuid
func (cli *ZSClient) GetBareMetal2ChassisPowerStatus(uuid string) (*view.GetBareMetal2ChassisPowerStatusView, error) {
	var resp view.GetBareMetal2ChassisPowerStatusView
	if err := cli.GetWithRespKey("v1/baremetal2/chassis", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBareMetal2ChassisState changes BareMetal2ChassisState
func (cli *ZSClient) ChangeBareMetal2ChassisState(uuid string, params param.ChangeBareMetal2ChassisStateParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/chassis", uuid, "", map[string]interface{}{
		"changeBareMetal2ChassisState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOnBareMetal2Chassis operates on PowerOnBareMetal2Chassis
func (cli *ZSClient) PowerOnBareMetal2Chassis(uuid string, params param.PowerOnBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/chassis", uuid, "", map[string]interface{}{
		"powerOnBareMetal2Chassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerResetBareMetal2Chassis operates on PowerResetBareMetal2Chassis
func (cli *ZSClient) PowerResetBareMetal2Chassis(uuid string, params param.PowerResetBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/chassis", uuid, "", map[string]interface{}{
		"powerResetBareMetal2Chassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// InspectBareMetal2ChassisByInstance operates on BareMetal2ChassisByInstance
func (cli *ZSClient) InspectBareMetal2ChassisByInstance(uuid string, params param.InspectBareMetal2ChassisByInstanceParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/bm-instances", uuid, "", map[string]interface{}{
		"inspectBareMetal2ChassisByInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
