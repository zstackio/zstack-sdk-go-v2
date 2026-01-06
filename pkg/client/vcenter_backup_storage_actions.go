// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVCenterBackupStorage queries VCenterBackupStorage list
func (cli *ZSClient) QueryVCenterBackupStorage(params *param.QueryParam) ([]view.VCenterBackupStorageInventoryView, error) {
	var resp []view.VCenterBackupStorageInventoryView
	return resp, cli.List("v1/vcenters/backup-storage", params, &resp)
}
