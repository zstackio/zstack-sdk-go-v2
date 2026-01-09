// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryZBoxBackup queries ZBoxBackup list
func (cli *ZSClient) QueryZBoxBackup(params *param.QueryParam) ([]view.ZBoxBackupInventoryView, error) {
	var resp []view.ZBoxBackupInventoryView
	return resp, cli.List("v1/externalbackup/zbox", params, &resp)
}

func (cli *ZSClient) GetZBoxBackup(uuid string) (*view.ZBoxBackupInventoryView, error) {
	var resp view.ZBoxBackupInventoryView
	if err := cli.Get("v1/externalbackup/zbox", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateZBoxBackup creates ZBoxBackup
func (cli *ZSClient) CreateZBoxBackup(params param.CreateZBoxBackupParam) (*view.ExternalBackupInventoryView, error) {
	var resp view.CreateExternalBackupEventView
	if err := cli.Post("v1/externalbackup/zbox", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// CreateZBoxBackupAsync Async
func (cli *ZSClient) CreateZBoxBackupAsync(params param.CreateZBoxBackupParam) (string, error) {

	resource := "v1/externalbackup/zbox"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
