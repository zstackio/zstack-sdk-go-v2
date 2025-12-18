// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryBuildAppExportHistory queries BuildAppExportHistory list
func (cli *ZSClient) QueryBuildAppExportHistory(params param.QueryParam) ([]view.BuildAppExportHistoryInventoryView, error) {
	var resp []view.BuildAppExportHistoryInventoryView
	return resp, cli.List("v1/appcenter/exportapp", &params, &resp)
}
