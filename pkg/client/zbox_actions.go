// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddZBox adds ZBox
func (cli *ZSClient) AddZBox(params param.AddZBoxParam) (*view.ZBoxInventoryView, error) {
	var resp view.AddZBoxEventView
	if err := cli.Post("v1/zbox", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryZBox queries ZBox list
func (cli *ZSClient) QueryZBox(params *param.QueryParam) ([]view.ZBoxInventoryView, error) {
	var resp []view.ZBoxInventoryView
	return resp, cli.List("v1/zbox", params, &resp)
}

func (cli *ZSClient) GetZBox(uuid string) (*view.ZBoxInventoryView, error) {
	var resp view.ZBoxInventoryView
	if err := cli.Get("v1/zbox", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
