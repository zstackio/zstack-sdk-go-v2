// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddIpRange adds IpRange
func (cli *ZSClient) AddIpRange(ctx context.Context, params param.AddIpRangeParam) (*view.IpRangeInventoryView, error) {
	resp := view.IpRangeInventoryView{}
	if err := cli.Post(ctx, "v1/l3-networks/{l3NetworkUuid}/ip-ranges", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateIpRange updates IpRange
func (cli *ZSClient) UpdateIpRange(ctx context.Context, uuid string, params param.UpdateIpRangeParam) (*view.IpRangeInventoryView, error) {
	resp := view.IpRangeInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/l3-networks/ip-ranges", uuid, "", map[string]interface{}{
		"updateIpRange": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryIpRange queries IpRange list
func (cli *ZSClient) QueryIpRange(ctx context.Context, params *param.QueryParam) ([]view.IpRangeInventoryView, error) {
	var resp []view.IpRangeInventoryView
	return resp, cli.List(ctx, "v1/l3-networks/ip-ranges", params, &resp)
}

func (cli *ZSClient) GetIpRange(ctx context.Context, uuid string) (*view.IpRangeInventoryView, error) {
	var resp view.IpRangeInventoryView
	if err := cli.Get(ctx, "v1/l3-networks/ip-ranges", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIpRange Pagination
func (cli *ZSClient) PageIpRange(ctx context.Context, params *param.QueryParam) ([]view.IpRangeInventoryView, int, error) {
	var ipRanges []view.IpRangeInventoryView
	total, err := cli.Page(ctx, "v1/l3-networks/ip-ranges", params, &ipRanges)
	return ipRanges, total, err
}
// DeleteIpRange deletes IpRange
func (cli *ZSClient) DeleteIpRange(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l3-networks/ip-ranges", uuid, string(deleteMode))
}
