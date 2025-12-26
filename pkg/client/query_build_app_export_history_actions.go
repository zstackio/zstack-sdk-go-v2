// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryBuildAppExportHistory queries BuildAppExportHistory list
func (cli *ZSClient) QueryBuildAppExportHistory(params *param.QueryParam) ([]view.BuildAppExportHistoryInventoryView, error) {
	var resp []view.BuildAppExportHistoryInventoryView
	return resp, cli.List("v1/appcenter/exportapp", params, &resp)
}
