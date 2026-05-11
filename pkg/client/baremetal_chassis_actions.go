// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBaremetalChassis queries BaremetalChassis list
func (cli *ZSClient) QueryBaremetalChassis(ctx context.Context, params *param.QueryParam) ([]view.BaremetalChassisInventoryView, error) {
	var resp []view.BaremetalChassisInventoryView
	return resp, cli.List(ctx, "v1/baremetal/chassis", params, &resp)
}

func (cli *ZSClient) GetBaremetalChassis(ctx context.Context, uuid string) (*view.BaremetalChassisInventoryView, error) {
	var resp view.BaremetalChassisInventoryView
	if err := cli.Get(ctx, "v1/baremetal/chassis", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBaremetalChassis Pagination
func (cli *ZSClient) PageBaremetalChassis(ctx context.Context, params *param.QueryParam) ([]view.BaremetalChassisInventoryView, int, error) {
	var baremetalChassis []view.BaremetalChassisInventoryView
	total, err := cli.Page(ctx, "v1/baremetal/chassis", params, &baremetalChassis)
	return baremetalChassis, total, err
}
// InspectBaremetalChassis operates on BaremetalChassis
func (cli *ZSClient) InspectBaremetalChassis(ctx context.Context, uuid string, params param.InspectBaremetalChassisParam) (*view.BaremetalChassisInventoryView, error) {
	resp := view.BaremetalChassisInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal/chassis", uuid, "", map[string]interface{}{
		"inspectBaremetalChassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateBaremetalChassis updates BaremetalChassis
func (cli *ZSClient) UpdateBaremetalChassis(ctx context.Context, uuid string, params param.UpdateBaremetalChassisParam) (*view.BaremetalChassisInventoryView, error) {
	resp := view.BaremetalChassisInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal/chassis", uuid, "", map[string]interface{}{
		"updateBaremetalChassis": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteBaremetalChassis deletes BaremetalChassis
func (cli *ZSClient) DeleteBaremetalChassis(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/baremetal/chassis", uuid, string(deleteMode))
}
// CreateBaremetalChassis creates BaremetalChassis
func (cli *ZSClient) CreateBaremetalChassis(ctx context.Context, params param.CreateBaremetalChassisParam) (*view.BaremetalChassisInventoryView, error) {
	resp := view.BaremetalChassisInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/baremetal/chassis", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
