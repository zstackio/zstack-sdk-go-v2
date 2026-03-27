// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// InspectBareMetal2Chassis operates on BareMetal2Chassis
func (cli *ZSClient) InspectBareMetal2Chassis(ctx context.Context, uuid string, params param.InspectBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/chassis", uuid, "", map[string]interface{}{
		"inspectBareMetal2Chassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateBareMetal2Chassis updates BareMetal2Chassis
func (cli *ZSClient) UpdateBareMetal2Chassis(ctx context.Context, uuid string, params param.UpdateBareMetal2ChassisParam) (*view.BareMetal2ChassisInventoryView, error) {
	resp := view.BareMetal2ChassisInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/chassis", uuid, "", map[string]interface{}{
		"updateBareMetal2Chassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryBareMetal2Chassis queries BareMetal2Chassis list
func (cli *ZSClient) QueryBareMetal2Chassis(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2ChassisInventoryView, error) {
	var resp []view.BareMetal2ChassisInventoryView
	return resp, cli.List(ctx, "v1/baremetal2/chassis", params, &resp)
}

func (cli *ZSClient) GetBareMetal2Chassis(ctx context.Context, uuid string) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.BareMetal2ChassisInventoryView
	if err := cli.Get(ctx, "v1/baremetal2/chassis", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBareMetal2Chassis Pagination
func (cli *ZSClient) PageBareMetal2Chassis(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2ChassisInventoryView, int, error) {
	var bareMetal2Chassis []view.BareMetal2ChassisInventoryView
	total, err := cli.Page(ctx, "v1/baremetal2/chassis", params, &bareMetal2Chassis)
	return bareMetal2Chassis, total, err
}
// DeleteBareMetal2Chassis deletes BareMetal2Chassis
func (cli *ZSClient) DeleteBareMetal2Chassis(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/baremetal2/chassis", uuid, string(deleteMode))
}
