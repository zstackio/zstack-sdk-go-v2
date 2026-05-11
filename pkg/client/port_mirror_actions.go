// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreatePortMirror creates PortMirror
func (cli *ZSClient) CreatePortMirror(ctx context.Context, params param.CreatePortMirrorParam) (*view.PortMirrorInventoryView, error) {
	resp := view.PortMirrorInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/port-mirrors", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryPortMirror queries PortMirror list
func (cli *ZSClient) QueryPortMirror(ctx context.Context, params *param.QueryParam) ([]view.PortMirrorInventoryView, error) {
	var resp []view.PortMirrorInventoryView
	return resp, cli.List(ctx, "v1/port-mirrors", params, &resp)
}

func (cli *ZSClient) GetPortMirror(ctx context.Context, uuid string) (*view.PortMirrorInventoryView, error) {
	var resp view.PortMirrorInventoryView
	if err := cli.Get(ctx, "v1/port-mirrors", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePortMirror Pagination
func (cli *ZSClient) PagePortMirror(ctx context.Context, params *param.QueryParam) ([]view.PortMirrorInventoryView, int, error) {
	var portMirrors []view.PortMirrorInventoryView
	total, err := cli.Page(ctx, "v1/port-mirrors", params, &portMirrors)
	return portMirrors, total, err
}
// DeletePortMirror deletes PortMirror
func (cli *ZSClient) DeletePortMirror(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/port-mirrors", uuid, string(deleteMode))
}
// UpdatePortMirror updates PortMirror
func (cli *ZSClient) UpdatePortMirror(ctx context.Context, uuid string, params param.UpdatePortMirrorParam) (*view.PortMirrorInventoryView, error) {
	resp := view.PortMirrorInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/port-mirrors", uuid, "", map[string]interface{}{
		"updatePortMirror": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
