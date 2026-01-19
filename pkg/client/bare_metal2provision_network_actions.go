// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateBareMetal2ProvisionNetwork updates BareMetal2ProvisionNetwork
func (cli *ZSClient) UpdateBareMetal2ProvisionNetwork(uuid string, params param.UpdateBareMetal2ProvisionNetworkParam) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	resp := view.BareMetal2ProvisionNetworkInventoryView{}
	if err := cli.Put("v1/baremetal2/provision-networks", uuid, map[string]interface{}{
		"updateBareMetal2ProvisionNetwork": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryBareMetal2ProvisionNetwork queries BareMetal2ProvisionNetwork list
func (cli *ZSClient) QueryBareMetal2ProvisionNetwork(params *param.QueryParam) ([]view.BareMetal2ProvisionNetworkInventoryView, error) {
	var resp []view.BareMetal2ProvisionNetworkInventoryView
	return resp, cli.List("v1/baremetal2/provision-networks", params, &resp)
}

func (cli *ZSClient) GetBareMetal2ProvisionNetwork(uuid string) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	var resp view.BareMetal2ProvisionNetworkInventoryView
	if err := cli.Get("v1/baremetal2/provision-networks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBareMetal2ProvisionNetwork Pagination
func (cli *ZSClient) PageBareMetal2ProvisionNetwork(params *param.QueryParam) ([]view.BareMetal2ProvisionNetworkInventoryView, int, error) {
	var bareMetal2ProvisionNetworks []view.BareMetal2ProvisionNetworkInventoryView
	total, err := cli.Page("v1/baremetal2/provision-networks", params, &bareMetal2ProvisionNetworks)
	return bareMetal2ProvisionNetworks, total, err
}
// CreateBareMetal2ProvisionNetwork creates BareMetal2ProvisionNetwork
func (cli *ZSClient) CreateBareMetal2ProvisionNetwork(params param.CreateBareMetal2ProvisionNetworkParam) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	resp := view.BareMetal2ProvisionNetworkInventoryView{}
	if err := cli.Post("v1/baremetal2/provision-networks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteBareMetal2ProvisionNetwork deletes BareMetal2ProvisionNetwork
func (cli *ZSClient) DeleteBareMetal2ProvisionNetwork(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/provision-networks", uuid, string(deleteMode))
}
