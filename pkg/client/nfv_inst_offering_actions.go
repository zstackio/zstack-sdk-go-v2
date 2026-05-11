// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNfvInstOffering queries NfvInstOffering list
func (cli *ZSClient) QueryNfvInstOffering(ctx context.Context, params *param.QueryParam) ([]view.NfvInstOfferingInventoryView, error) {
	var resp []view.NfvInstOfferingInventoryView
	return resp, cli.List(ctx, "v1/instance-offerings/nfvinst", params, &resp)
}

func (cli *ZSClient) GetNfvInstOffering(ctx context.Context, uuid string) (*view.NfvInstOfferingInventoryView, error) {
	var resp view.NfvInstOfferingInventoryView
	if err := cli.Get(ctx, "v1/instance-offerings/nfvinst", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageNfvInstOffering Pagination
func (cli *ZSClient) PageNfvInstOffering(ctx context.Context, params *param.QueryParam) ([]view.NfvInstOfferingInventoryView, int, error) {
	var nfvInstOfferings []view.NfvInstOfferingInventoryView
	total, err := cli.Page(ctx, "v1/instance-offerings/nfvinst", params, &nfvInstOfferings)
	return nfvInstOfferings, total, err
}
// CreateNfvInstOffering creates NfvInstOffering
func (cli *ZSClient) CreateNfvInstOffering(ctx context.Context, params param.CreateNfvInstOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/instance-offerings/nfvinst", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
