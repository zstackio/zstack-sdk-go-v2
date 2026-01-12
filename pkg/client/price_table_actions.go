// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreatePriceTable creates PriceTable
func (cli *ZSClient) CreatePriceTable(params param.CreatePriceTableParam) (*view.PriceTableInventoryView, error) {
	var resp view.CreatePriceTableEventView
	if err := cli.Post("v1/billings/price-tables", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdatePriceTable updates PriceTable
func (cli *ZSClient) UpdatePriceTable(uuid string, params param.UpdatePriceTableParam) (*view.PriceTableInventoryView, error) {
	var resp view.UpdatePriceTableEventView
	err := cli.PutWithSpec("v1/billings/price-tables", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeletePriceTable deletes PriceTable
func (cli *ZSClient) DeletePriceTable(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/billings/price-tables", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// QueryPriceTable queries PriceTable list
func (cli *ZSClient) QueryPriceTable(params *param.QueryParam) ([]view.PriceTableInventoryView, error) {
	var resp []view.PriceTableInventoryView
	return resp, cli.List("v1/billings/price-tables", params, &resp)
}

func (cli *ZSClient) GetPriceTable(uuid string) (*view.PriceTableInventoryView, error) {
	var resp view.PriceTableInventoryView
	if err := cli.Get("v1/billings/price-tables", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
