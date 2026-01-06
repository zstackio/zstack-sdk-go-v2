// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteBuildAppExportHistory deletes BuildAppExportHistory
func (cli *ZSClient) DeleteBuildAppExportHistory(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/buildapp/exports/{buildAppUuid}", uuid, string(deleteMode))
}
// QueryBuildAppExportHistory queries BuildAppExportHistory list
func (cli *ZSClient) QueryBuildAppExportHistory(params *param.QueryParam) ([]view.BuildAppExportHistoryInventoryView, error) {
	var resp []view.BuildAppExportHistoryInventoryView
	return resp, cli.List("v1/appcenter/exportapp", params, &resp)
}
