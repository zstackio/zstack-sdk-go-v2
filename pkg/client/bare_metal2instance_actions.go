// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBareMetal2Instance queries BareMetal2Instance list
func (cli *ZSClient) QueryBareMetal2Instance(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2InstanceInventoryView, error) {
	var resp []view.BareMetal2InstanceInventoryView
	return resp, cli.List(ctx, "v1/baremetal2/bm-instances", params, &resp)
}

func (cli *ZSClient) GetBareMetal2Instance(ctx context.Context, uuid string) (*view.BareMetal2InstanceInventoryView, error) {
	var resp view.BareMetal2InstanceInventoryView
	if err := cli.Get(ctx, "v1/baremetal2/bm-instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBareMetal2Instance Pagination
func (cli *ZSClient) PageBareMetal2Instance(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2InstanceInventoryView, int, error) {
	var bareMetal2Instances []view.BareMetal2InstanceInventoryView
	total, err := cli.Page(ctx, "v1/baremetal2/bm-instances", params, &bareMetal2Instances)
	return bareMetal2Instances, total, err
}
// CreateBareMetal2Instance creates BareMetal2Instance
func (cli *ZSClient) CreateBareMetal2Instance(ctx context.Context, params param.CreateBareMetal2InstanceParam) (*view.BareMetal2InstanceInventoryView, error) {
	resp := view.BareMetal2InstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/baremetal2/bm-instances", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateBareMetal2InstanceAsync Async
func (cli *ZSClient) CreateBareMetal2InstanceAsync(ctx context.Context, params param.CreateBareMetal2InstanceParam) (string, error) {

	resource := "v1/baremetal2/bm-instances"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
// ReconnectBareMetal2Instance operates on BareMetal2Instance
func (cli *ZSClient) ReconnectBareMetal2Instance(ctx context.Context, uuid string, params param.ReconnectBareMetal2InstanceParam) (*view.BareMetal2InstanceInventoryView, error) {
	resp := view.BareMetal2InstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/bm-instances", uuid, "", map[string]interface{}{
		"reconnectBareMetal2Instance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StartBareMetal2Instance starts BareMetal2Instance
func (cli *ZSClient) StartBareMetal2Instance(ctx context.Context, uuid string, params param.StartBareMetal2InstanceParam) (*view.BareMetal2InstanceInventoryView, error) {
	resp := view.BareMetal2InstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/bm-instances", uuid, "", map[string]interface{}{
		"startBareMetal2Instance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateBareMetal2Instance updates BareMetal2Instance
func (cli *ZSClient) UpdateBareMetal2Instance(ctx context.Context, uuid string, params param.UpdateBareMetal2InstanceParam) (*view.BareMetal2InstanceInventoryView, error) {
	resp := view.BareMetal2InstanceInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/baremetal2/bm-instances", uuid, "", map[string]interface{}{
		"updateBareMetal2Instance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
