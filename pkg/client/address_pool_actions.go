// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAddressPool queries AddressPool list
func (cli *ZSClient) QueryAddressPool(params *param.QueryParam) ([]view.AddressPoolInventoryView, error) {
	var resp []view.AddressPoolInventoryView
	return resp, cli.List("v1/l3-networks/address-pools", params, &resp)
}

func (cli *ZSClient) GetAddressPool(uuid string) (*view.AddressPoolInventoryView, error) {
	var resp view.AddressPoolInventoryView
	if err := cli.Get("v1/l3-networks/address-pools", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAddressPool Pagination
func (cli *ZSClient) PageAddressPool(params *param.QueryParam) ([]view.AddressPoolInventoryView, int, error) {
	var addressPools []view.AddressPoolInventoryView
	total, err := cli.Page("v1/l3-networks/address-pools", params, &addressPools)
	return addressPools, total, err
}
