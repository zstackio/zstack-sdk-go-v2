// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateLogServer updates LogServer
func (cli *ZSClient) UpdateLogServer(ctx context.Context, params param.UpdateLogServerParam) (*view.LogServerInventoryView, error) {
	resp := view.LogServerInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/log/servers", "", "", map[string]interface{}{
		"updateLogServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteLogServer deletes LogServer
func (cli *ZSClient) DeleteLogServer(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/log/servers", uuid, string(deleteMode))
}
// QueryLogServer queries LogServer list
func (cli *ZSClient) QueryLogServer(ctx context.Context, params *param.QueryParam) ([]view.LogServerInventoryView, error) {
	var resp []view.LogServerInventoryView
	return resp, cli.List(ctx, "v1/log/servers", params, &resp)
}

func (cli *ZSClient) GetLogServer(ctx context.Context, uuid string) (*view.LogServerInventoryView, error) {
	var resp view.LogServerInventoryView
	if err := cli.Get(ctx, "v1/log/servers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLogServer Pagination
func (cli *ZSClient) PageLogServer(ctx context.Context, params *param.QueryParam) ([]view.LogServerInventoryView, int, error) {
	var logServers []view.LogServerInventoryView
	total, err := cli.Page(ctx, "v1/log/servers", params, &logServers)
	return logServers, total, err
}
// AddLogServer adds LogServer
func (cli *ZSClient) AddLogServer(ctx context.Context, params param.AddLogServerParam) (*view.LogServerInventoryView, error) {
	resp := view.LogServerInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/log/servers", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
