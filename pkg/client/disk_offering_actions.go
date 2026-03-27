// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteDiskOffering deletes DiskOffering
func (cli *ZSClient) DeleteDiskOffering(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/disk-offerings", uuid, string(deleteMode))
}
// CreateDiskOffering creates DiskOffering
func (cli *ZSClient) CreateDiskOffering(ctx context.Context, params param.CreateDiskOfferingParam) (*view.DiskOfferingInventoryView, error) {
	resp := view.DiskOfferingInventoryView{}
	if err := cli.Post(ctx, "v1/disk-offerings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateDiskOffering updates DiskOffering
func (cli *ZSClient) UpdateDiskOffering(ctx context.Context, uuid string, params param.UpdateDiskOfferingParam) (*view.DiskOfferingInventoryView, error) {
	resp := view.DiskOfferingInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/disk-offerings", uuid, "", map[string]interface{}{
		"updateDiskOffering": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryDiskOffering queries DiskOffering list
func (cli *ZSClient) QueryDiskOffering(ctx context.Context, params *param.QueryParam) ([]view.DiskOfferingInventoryView, error) {
	var resp []view.DiskOfferingInventoryView
	return resp, cli.List(ctx, "v1/disk-offerings", params, &resp)
}

func (cli *ZSClient) GetDiskOffering(ctx context.Context, uuid string) (*view.DiskOfferingInventoryView, error) {
	var resp view.DiskOfferingInventoryView
	if err := cli.Get(ctx, "v1/disk-offerings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageDiskOffering Pagination
func (cli *ZSClient) PageDiskOffering(ctx context.Context, params *param.QueryParam) ([]view.DiskOfferingInventoryView, int, error) {
	var diskOfferings []view.DiskOfferingInventoryView
	total, err := cli.Page(ctx, "v1/disk-offerings", params, &diskOfferings)
	return diskOfferings, total, err
}
