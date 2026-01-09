// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteExternalBackup deletes ExternalBackup
func (cli *ZSClient) DeleteExternalBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/externalbackup", uuid, string(deleteMode))
}
// QueryExternalBackup queries ExternalBackup list
func (cli *ZSClient) QueryExternalBackup(params *param.QueryParam) ([]view.ExternalBackupInventoryView, error) {
	var resp []view.ExternalBackupInventoryView
	return resp, cli.List("v1/externalbackup", params, &resp)
}

func (cli *ZSClient) GetExternalBackup(uuid string) (*view.ExternalBackupInventoryView, error) {
	var resp view.ExternalBackupInventoryView
	if err := cli.Get("v1/externalbackup", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
