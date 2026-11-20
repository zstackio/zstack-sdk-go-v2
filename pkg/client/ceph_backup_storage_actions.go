// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryCephBackupStorage queries CephBackupStorage list
func (cli *ZSClient) QueryCephBackupStorage(params *param.QueryParam) ([]view.BackupStorageInventoryView, error) {
	var resp []view.BackupStorageInventoryView
	return resp, cli.List("v1/backup-storage/ceph", params, &resp)
}

func (cli *ZSClient) GetCephBackupStorage(uuid string) (*view.BackupStorageInventoryView, error) {
	var resp view.BackupStorageInventoryView
	if err := cli.Get("v1/backup-storage/ceph", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCephBackupStorage Pagination
func (cli *ZSClient) PageCephBackupStorage(params *param.QueryParam) ([]view.BackupStorageInventoryView, int, error) {
	var cephBackupStorages []view.BackupStorageInventoryView
	total, err := cli.Page("v1/backup-storage/ceph", params, &cephBackupStorages)
	return cephBackupStorages, total, err
}
// AddCephBackupStorage adds CephBackupStorage
func (cli *ZSClient) AddCephBackupStorage(params param.AddCephBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.Post("v1/backup-storage/ceph", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddMonToCephBackupStorage adds MonToCephBackupStorage
func (cli *ZSClient) AddMonToCephBackupStorage(uuid string, params param.AddMonToCephBackupStorageParam) (*view.CephBackupStorageInventoryView, error) {
	resp := view.CephBackupStorageInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/backup-storage/ceph/%s/mons", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveMonFromCephBackupStorage requires additional arguments.
//
// Deprecated: use RemoveMonFromCephBackupStorageWithParams to supply the required fields.
func (cli *ZSClient) RemoveMonFromCephBackupStorage(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveMonFromCephBackupStorage requires additional arguments; use RemoveMonFromCephBackupStorageWithParams")
}

// RemoveMonFromCephBackupStorageWithParams sends the required route and query parameters for RemoveMonFromCephBackupStorage.
func (cli *ZSClient) RemoveMonFromCephBackupStorageWithParams(uuid string, deleteMode param.DeleteMode, params param.RemoveMonFromCephBackupStorageParamDetail) error {
	query := url.Values{
		"monHostnames": params.MonHostnames,
	}
	return cli.deleteActionWithQuery("v1/backup-storage/ceph", uuid, "mons", deleteMode, query, "monHostnames")
}
