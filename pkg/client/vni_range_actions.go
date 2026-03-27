// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVniRange updates VniRange
func (cli *ZSClient) UpdateVniRange(ctx context.Context, uuid string, params param.UpdateVniRangeParam) (*view.VniRangeInventoryView, error) {
	resp := view.VniRangeInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/l2-networks/vxlan-pool/vni-ranges", uuid, "", map[string]interface{}{
		"updateVniRange": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVniRange queries VniRange list
func (cli *ZSClient) QueryVniRange(ctx context.Context, params *param.QueryParam) ([]view.VniRangeInventoryView, error) {
	var resp []view.VniRangeInventoryView
	return resp, cli.List(ctx, "v1/l2-networks/vxlan-pool/vni-range", params, &resp)
}

func (cli *ZSClient) GetVniRange(ctx context.Context, uuid string) (*view.VniRangeInventoryView, error) {
	var resp view.VniRangeInventoryView
	if err := cli.Get(ctx, "v1/l2-networks/vxlan-pool/vni-range", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVniRange Pagination
func (cli *ZSClient) PageVniRange(ctx context.Context, params *param.QueryParam) ([]view.VniRangeInventoryView, int, error) {
	var vniRanges []view.VniRangeInventoryView
	total, err := cli.Page(ctx, "v1/l2-networks/vxlan-pool/vni-range", params, &vniRanges)
	return vniRanges, total, err
}
// CreateVniRange creates VniRange
func (cli *ZSClient) CreateVniRange(ctx context.Context, params param.CreateVniRangeParam) (*view.VniRangeInventoryView, error) {
	resp := view.VniRangeInventoryView{}
	if err := cli.Post(ctx, "v1/l2-networks/vxlan-pool/{l2NetworkUuid}/vni-ranges", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteVniRange deletes VniRange
func (cli *ZSClient) DeleteVniRange(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l2-networks/vxlan-pool/vni-ranges", uuid, string(deleteMode))
}
