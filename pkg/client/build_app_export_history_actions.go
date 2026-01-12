// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteBuildAppExportHistory deletes BuildAppExportHistory
func (cli *ZSClient) DeleteBuildAppExportHistory(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/buildapp/exports", uuid, string(deleteMode))
}
// QueryBuildAppExportHistory queries BuildAppExportHistory list
func (cli *ZSClient) QueryBuildAppExportHistory(params *param.QueryParam) ([]view.BuildAppExportHistoryInventoryView, error) {
	var resp []view.BuildAppExportHistoryInventoryView
	return resp, cli.List("v1/appcenter/exportapp", params, &resp)
}

func (cli *ZSClient) GetBuildAppExportHistory(uuid string) (*view.BuildAppExportHistoryInventoryView, error) {
	var resp view.BuildAppExportHistoryInventoryView
	if err := cli.Get("v1/appcenter/exportapp", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
