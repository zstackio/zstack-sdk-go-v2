// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSlbOffering creates SlbOffering
func (cli *ZSClient) CreateSlbOffering(ctx context.Context, params param.CreateSlbOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.Post(ctx, "v1/instance-offerings/slb", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySlbOffering queries SlbOffering list
func (cli *ZSClient) QuerySlbOffering(ctx context.Context, params *param.QueryParam) ([]view.SlbOfferingInventoryView, error) {
	var resp []view.SlbOfferingInventoryView
	return resp, cli.List(ctx, "v1/instance-offerings/slb", params, &resp)
}

func (cli *ZSClient) GetSlbOffering(ctx context.Context, uuid string) (*view.SlbOfferingInventoryView, error) {
	var resp view.SlbOfferingInventoryView
	if err := cli.Get(ctx, "v1/instance-offerings/slb", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSlbOffering Pagination
func (cli *ZSClient) PageSlbOffering(ctx context.Context, params *param.QueryParam) ([]view.SlbOfferingInventoryView, int, error) {
	var slbOfferings []view.SlbOfferingInventoryView
	total, err := cli.Page(ctx, "v1/instance-offerings/slb", params, &slbOfferings)
	return slbOfferings, total, err
}
