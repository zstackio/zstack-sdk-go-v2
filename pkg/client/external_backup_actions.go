// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteExternalBackup deletes ExternalBackup
func (cli *ZSClient) DeleteExternalBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/externalbackup/{uuid}", uuid, string(deleteMode))
}
// QueryExternalBackup queries ExternalBackup list
func (cli *ZSClient) QueryExternalBackup(params *param.QueryParam) ([]view.ExternalBackupInventoryView, error) {
	var resp []view.ExternalBackupInventoryView
	return resp, cli.List("v1/externalbackup", params, &resp)
}
