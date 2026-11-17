// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// UpdateResourcePrice updates ResourcePrice
func (cli *ZSClient) UpdateResourcePrice(uuid string, params param.UpdateResourcePriceParam) (*view.PriceInventoryView, error) {
	resp := view.PriceInventoryView{}
	if err := cli.PutWithRespKey("v1/billings/prices", uuid, "", map[string]interface{}{
		"updateResourcePrice": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteResourcePrice deletes ResourcePrice
func (cli *ZSClient) DeleteResourcePrice(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/billings/prices", uuid, string(deleteMode))
}

// QueryResourcePrice queries ResourcePrice list
func (cli *ZSClient) QueryResourcePrice(params *param.QueryParam) ([]view.PriceInventoryView, error) {
	var resp []view.PriceInventoryView
	return resp, cli.List("v1/billings/prices", params, &resp)
}

func (cli *ZSClient) GetResourcePrice(uuid string) (*view.PriceInventoryView, error) {
	var resp view.PriceInventoryView
	if err := cli.Get("v1/billings/prices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageResourcePrice Pagination
func (cli *ZSClient) PageResourcePrice(params *param.QueryParam) ([]view.PriceInventoryView, int, error) {
	var resourcePrices []view.PriceInventoryView
	total, err := cli.Page("v1/billings/prices", params, &resourcePrices)
	return resourcePrices, total, err
}

// CreateResourcePrice creates ResourcePrice
func (cli *ZSClient) CreateResourcePrice(params param.CreateResourcePriceParam) (*view.PriceInventoryView, error) {
	resp := view.PriceInventoryView{}
	if err := cli.Post("v1/billings/prices", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
