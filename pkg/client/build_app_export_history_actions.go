// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteBuildAppExportHistory deletes BuildAppExportHistory
func (cli *ZSClient) DeleteBuildAppExportHistory(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/buildapp/exports", uuid, string(deleteMode))
}
// QueryBuildAppExportHistory queries BuildAppExportHistory list
func (cli *ZSClient) QueryBuildAppExportHistory(ctx context.Context, params *param.QueryParam) ([]view.BuildAppExportHistoryInventoryView, error) {
	var resp []view.BuildAppExportHistoryInventoryView
	return resp, cli.List(ctx, "v1/appcenter/exportapp", params, &resp)
}

func (cli *ZSClient) GetBuildAppExportHistory(ctx context.Context, uuid string) (*view.BuildAppExportHistoryInventoryView, error) {
	var resp view.BuildAppExportHistoryInventoryView
	if err := cli.Get(ctx, "v1/appcenter/exportapp", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBuildAppExportHistory Pagination
func (cli *ZSClient) PageBuildAppExportHistory(ctx context.Context, params *param.QueryParam) ([]view.BuildAppExportHistoryInventoryView, int, error) {
	var buildAppExportHistories []view.BuildAppExportHistoryInventoryView
	total, err := cli.Page(ctx, "v1/appcenter/exportapp", params, &buildAppExportHistories)
	return buildAppExportHistories, total, err
}
