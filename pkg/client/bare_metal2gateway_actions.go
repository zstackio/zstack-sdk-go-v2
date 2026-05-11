// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ReconnectBareMetal2Gateway operates on BareMetal2Gateway
func (cli *ZSClient) ReconnectBareMetal2Gateway(ctx context.Context, uuid string, params param.ReconnectBareMetal2GatewayParam) (*view.BareMetal2GatewayInventoryView, error) {
	resp := view.BareMetal2GatewayInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/gateways", uuid, "", map[string]interface{}{
		"reconnectBareMetal2Gateway": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateBareMetal2Gateway updates BareMetal2Gateway
func (cli *ZSClient) UpdateBareMetal2Gateway(ctx context.Context, uuid string, params param.UpdateBareMetal2GatewayParam) (*view.BareMetal2GatewayInventoryView, error) {
	resp := view.BareMetal2GatewayInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/gateways", uuid, "", map[string]interface{}{
		"updateBareMetal2Gateway": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryBareMetal2Gateway queries BareMetal2Gateway list
func (cli *ZSClient) QueryBareMetal2Gateway(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2GatewayInventoryView, error) {
	var resp []view.BareMetal2GatewayInventoryView
	return resp, cli.List(ctx, "v1/baremetal2/gateways", params, &resp)
}

func (cli *ZSClient) GetBareMetal2Gateway(ctx context.Context, uuid string) (*view.BareMetal2GatewayInventoryView, error) {
	var resp view.BareMetal2GatewayInventoryView
	if err := cli.Get(ctx, "v1/baremetal2/gateways", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBareMetal2Gateway Pagination
func (cli *ZSClient) PageBareMetal2Gateway(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2GatewayInventoryView, int, error) {
	var bareMetal2Gateways []view.BareMetal2GatewayInventoryView
	total, err := cli.Page(ctx, "v1/baremetal2/gateways", params, &bareMetal2Gateways)
	return bareMetal2Gateways, total, err
}
// DeleteBareMetal2Gateway deletes BareMetal2Gateway
func (cli *ZSClient) DeleteBareMetal2Gateway(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/baremetal2/gateways", uuid, string(deleteMode))
}
// AddBareMetal2Gateway adds BareMetal2Gateway
func (cli *ZSClient) AddBareMetal2Gateway(ctx context.Context, params param.AddBareMetal2GatewayParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/baremetal2/gateways", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
