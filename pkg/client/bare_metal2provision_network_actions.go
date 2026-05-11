// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateBareMetal2ProvisionNetwork updates BareMetal2ProvisionNetwork
func (cli *ZSClient) UpdateBareMetal2ProvisionNetwork(ctx context.Context, uuid string, params param.UpdateBareMetal2ProvisionNetworkParam) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	resp := view.BareMetal2ProvisionNetworkInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/provision-networks", uuid, "", map[string]interface{}{
		"updateBareMetal2ProvisionNetwork": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryBareMetal2ProvisionNetwork queries BareMetal2ProvisionNetwork list
func (cli *ZSClient) QueryBareMetal2ProvisionNetwork(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2ProvisionNetworkInventoryView, error) {
	var resp []view.BareMetal2ProvisionNetworkInventoryView
	return resp, cli.List(ctx, "v1/baremetal2/provision-networks", params, &resp)
}

func (cli *ZSClient) GetBareMetal2ProvisionNetwork(ctx context.Context, uuid string) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	var resp view.BareMetal2ProvisionNetworkInventoryView
	if err := cli.Get(ctx, "v1/baremetal2/provision-networks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBareMetal2ProvisionNetwork Pagination
func (cli *ZSClient) PageBareMetal2ProvisionNetwork(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2ProvisionNetworkInventoryView, int, error) {
	var bareMetal2ProvisionNetworks []view.BareMetal2ProvisionNetworkInventoryView
	total, err := cli.Page(ctx, "v1/baremetal2/provision-networks", params, &bareMetal2ProvisionNetworks)
	return bareMetal2ProvisionNetworks, total, err
}
// CreateBareMetal2ProvisionNetwork creates BareMetal2ProvisionNetwork
func (cli *ZSClient) CreateBareMetal2ProvisionNetwork(ctx context.Context, params param.CreateBareMetal2ProvisionNetworkParam) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	resp := view.BareMetal2ProvisionNetworkInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/baremetal2/provision-networks", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteBareMetal2ProvisionNetwork deletes BareMetal2ProvisionNetwork
func (cli *ZSClient) DeleteBareMetal2ProvisionNetwork(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/baremetal2/provision-networks", uuid, string(deleteMode))
}
