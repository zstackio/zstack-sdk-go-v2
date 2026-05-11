// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreatePriceTable creates PriceTable
func (cli *ZSClient) CreatePriceTable(ctx context.Context, params param.CreatePriceTableParam) (*view.PriceTableInventoryView, error) {
	resp := view.PriceTableInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/billings/price-tables", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdatePriceTable updates PriceTable
func (cli *ZSClient) UpdatePriceTable(ctx context.Context, uuid string, params param.UpdatePriceTableParam) (*view.PriceTableInventoryView, error) {
	resp := view.PriceTableInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/billings/price-tables", uuid, "", map[string]interface{}{
		"updatePriceTable": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePriceTable deletes PriceTable
func (cli *ZSClient) DeletePriceTable(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/billings/price-tables", uuid, string(deleteMode))
}
// QueryPriceTable queries PriceTable list
func (cli *ZSClient) QueryPriceTable(ctx context.Context, params *param.QueryParam) ([]view.PriceTableInventoryView, error) {
	var resp []view.PriceTableInventoryView
	return resp, cli.List(ctx, "v1/billings/price-tables", params, &resp)
}

func (cli *ZSClient) GetPriceTable(ctx context.Context, uuid string) (*view.PriceTableInventoryView, error) {
	var resp view.PriceTableInventoryView
	if err := cli.Get(ctx, "v1/billings/price-tables", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePriceTable Pagination
func (cli *ZSClient) PagePriceTable(ctx context.Context, params *param.QueryParam) ([]view.PriceTableInventoryView, int, error) {
	var priceTables []view.PriceTableInventoryView
	total, err := cli.Page(ctx, "v1/billings/price-tables", params, &priceTables)
	return priceTables, total, err
}
