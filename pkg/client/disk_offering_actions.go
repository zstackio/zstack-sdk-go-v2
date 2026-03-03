// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteDiskOffering deletes DiskOffering
func (cli *ZSClient) DeleteDiskOffering(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/disk-offerings", uuid, string(deleteMode))
}
// CreateDiskOffering creates DiskOffering
func (cli *ZSClient) CreateDiskOffering(params param.CreateDiskOfferingParam) (*view.DiskOfferingInventoryView, error) {
	resp := view.DiskOfferingInventoryView{}
	if err := cli.Post("v1/disk-offerings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateDiskOffering updates DiskOffering
func (cli *ZSClient) UpdateDiskOffering(uuid string, params param.UpdateDiskOfferingParam) (*view.DiskOfferingInventoryView, error) {
	resp := view.DiskOfferingInventoryView{}
	if err := cli.PutWithRespKey("v1/disk-offerings", uuid, "", map[string]interface{}{
		"updateDiskOffering": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryDiskOffering queries DiskOffering list
func (cli *ZSClient) QueryDiskOffering(params *param.QueryParam) ([]view.DiskOfferingInventoryView, error) {
	var resp []view.DiskOfferingInventoryView
	return resp, cli.List("v1/disk-offerings", params, &resp)
}

func (cli *ZSClient) GetDiskOffering(uuid string) (*view.DiskOfferingInventoryView, error) {
	var resp view.DiskOfferingInventoryView
	if err := cli.Get("v1/disk-offerings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageDiskOffering Pagination
func (cli *ZSClient) PageDiskOffering(params *param.QueryParam) ([]view.DiskOfferingInventoryView, int, error) {
	var diskOfferings []view.DiskOfferingInventoryView
	total, err := cli.Page("v1/disk-offerings", params, &diskOfferings)
	return diskOfferings, total, err
}
