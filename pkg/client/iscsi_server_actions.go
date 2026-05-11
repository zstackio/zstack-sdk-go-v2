// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddIscsiServer adds IscsiServer
func (cli *ZSClient) AddIscsiServer(ctx context.Context, params param.AddIscsiServerParam) (*view.IscsiServerInventoryView, error) {
	resp := view.IscsiServerInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/storage-devices/iscsi/servers", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryIscsiServer queries IscsiServer list
func (cli *ZSClient) QueryIscsiServer(ctx context.Context, params *param.QueryParam) ([]view.IscsiServerInventoryView, error) {
	var resp []view.IscsiServerInventoryView
	return resp, cli.List(ctx, "v1/storage-devices/iscsi/servers", params, &resp)
}

func (cli *ZSClient) GetIscsiServer(ctx context.Context, uuid string) (*view.IscsiServerInventoryView, error) {
	var resp view.IscsiServerInventoryView
	if err := cli.Get(ctx, "v1/storage-devices/iscsi/servers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIscsiServer Pagination
func (cli *ZSClient) PageIscsiServer(ctx context.Context, params *param.QueryParam) ([]view.IscsiServerInventoryView, int, error) {
	var iscsiServers []view.IscsiServerInventoryView
	total, err := cli.Page(ctx, "v1/storage-devices/iscsi/servers", params, &iscsiServers)
	return iscsiServers, total, err
}
// DeleteIscsiServer deletes IscsiServer
func (cli *ZSClient) DeleteIscsiServer(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/storage-devices/iscsi/servers", uuid, string(deleteMode))
}
// RefreshIscsiServer operates on IscsiServer
func (cli *ZSClient) RefreshIscsiServer(ctx context.Context, uuid string, params param.RefreshIscsiServerParam) (*view.IscsiServerInventoryView, error) {
	resp := view.IscsiServerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/storage-devices/iscsi/servers/%s", uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateIscsiServer updates IscsiServer
func (cli *ZSClient) UpdateIscsiServer(ctx context.Context, uuid string, params param.UpdateIscsiServerParam) (*view.IscsiServerInventoryView, error) {
	resp := view.IscsiServerInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/storage-devices/iscsi/servers", uuid, "", map[string]interface{}{
		"updateIscsiServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
