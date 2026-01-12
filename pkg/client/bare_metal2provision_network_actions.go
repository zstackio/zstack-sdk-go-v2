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
	var resp view.UpdateBareMetal2ProvisionNetworkEventView
	if err := cli.Put("v1/baremetal2/provision-networks", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
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
// CreateBareMetal2ProvisionNetwork creates BareMetal2ProvisionNetwork
func (cli *ZSClient) CreateBareMetal2ProvisionNetwork(params param.CreateBareMetal2ProvisionNetworkParam) (*view.BareMetal2ProvisionNetworkInventoryView, error) {
	var resp view.CreateBareMetal2ProvisionNetworkEventView
	if err := cli.Post("v1/baremetal2/provision-networks", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteBareMetal2ProvisionNetwork deletes BareMetal2ProvisionNetwork
func (cli *ZSClient) DeleteBareMetal2ProvisionNetwork(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal2/provision-networks", uuid, string(deleteMode))
}
