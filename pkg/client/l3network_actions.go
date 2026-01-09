// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryL3Network queries L3Network list
func (cli *ZSClient) QueryL3Network(params *param.QueryParam) ([]view.L3NetworkInventoryView, error) {
	var resp []view.L3NetworkInventoryView
	return resp, cli.List("v1/l3-networks", params, &resp)
}

func (cli *ZSClient) GetL3Network(uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.Get("v1/l3-networks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateL3Network updates L3Network
func (cli *ZSClient) UpdateL3Network(uuid string, params param.UpdateL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	var resp view.UpdateL3NetworkEventView
	if err := cli.Put("v1/l3-networks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateL3Network creates L3Network
func (cli *ZSClient) CreateL3Network(params param.CreateL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	var resp view.CreateL3NetworkEventView
	if err := cli.Post("v1/l3-networks", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteL3Network deletes L3Network
func (cli *ZSClient) DeleteL3Network(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/l3-networks", uuid, string(deleteMode))
}
