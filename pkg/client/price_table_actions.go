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
	resp := view.PriceTableInventoryView{}
	if err := cli.Post("v1/billings/price-tables", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdatePriceTable updates PriceTable
func (cli *ZSClient) UpdatePriceTable(uuid string, params param.UpdatePriceTableParam) (*view.PriceTableInventoryView, error) {
	resp := view.PriceTableInventoryView{}
	if err := cli.PutWithRespKey("v1/billings/price-tables", uuid, "", map[string]interface{}{
		"updatePriceTable": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePriceTable deletes PriceTable
func (cli *ZSClient) DeletePriceTable(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/price-tables", uuid, string(deleteMode))
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

// PagePriceTable Pagination
func (cli *ZSClient) PagePriceTable(params *param.QueryParam) ([]view.PriceTableInventoryView, int, error) {
	var priceTables []view.PriceTableInventoryView
	total, err := cli.Page("v1/billings/price-tables", params, &priceTables)
	return priceTables, total, err
}
