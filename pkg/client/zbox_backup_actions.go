// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryZBoxBackup queries ZBoxBackup list
func (cli *ZSClient) QueryZBoxBackup(params *param.QueryParam) ([]view.ZBoxBackupInventoryView, error) {
	var resp []view.ZBoxBackupInventoryView
	return resp, cli.List("v1/externalbackup/zbox", params, &resp)
}
// CreateZBoxBackup creates ZBoxBackup
func (cli *ZSClient) CreateZBoxBackup(params param.CreateZBoxBackupParam) (*view.ExternalBackupInventoryView, error) {
	var resp view.CreateExternalBackupEventView
	if err := cli.Post("v1/externalbackup/zbox", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
