// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryDRSVmMigrationActivity queries DRSVmMigrationActivity list
func (cli *ZSClient) QueryDRSVmMigrationActivity(params *param.QueryParam) ([]view.DRSVmMigrationActivityInventoryView, error) {
	var resp []view.DRSVmMigrationActivityInventoryView
	return resp, cli.List("v1/clusters/drs/vm-migration-activities", params, &resp)
}
