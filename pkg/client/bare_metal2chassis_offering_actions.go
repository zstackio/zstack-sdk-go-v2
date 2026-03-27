// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateBareMetal2ChassisOffering updates BareMetal2ChassisOffering
func (cli *ZSClient) UpdateBareMetal2ChassisOffering(ctx context.Context, uuid string, params param.UpdateBareMetal2ChassisOfferingParam) (*view.BareMetal2ChassisOfferingInventoryView, error) {
	resp := view.BareMetal2ChassisOfferingInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/chassis/offerings", uuid, "", map[string]interface{}{
		"updateBareMetal2ChassisOffering": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryBareMetal2ChassisOffering queries BareMetal2ChassisOffering list
func (cli *ZSClient) QueryBareMetal2ChassisOffering(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2ChassisOfferingInventoryView, error) {
	var resp []view.BareMetal2ChassisOfferingInventoryView
	return resp, cli.List(ctx, "v1/baremetal2/chassis/offerings", params, &resp)
}

func (cli *ZSClient) GetBareMetal2ChassisOffering(ctx context.Context, uuid string) (*view.BareMetal2ChassisOfferingInventoryView, error) {
	var resp view.BareMetal2ChassisOfferingInventoryView
	if err := cli.Get(ctx, "v1/baremetal2/chassis/offerings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBareMetal2ChassisOffering Pagination
func (cli *ZSClient) PageBareMetal2ChassisOffering(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2ChassisOfferingInventoryView, int, error) {
	var bareMetal2ChassisOfferings []view.BareMetal2ChassisOfferingInventoryView
	total, err := cli.Page(ctx, "v1/baremetal2/chassis/offerings", params, &bareMetal2ChassisOfferings)
	return bareMetal2ChassisOfferings, total, err
}
