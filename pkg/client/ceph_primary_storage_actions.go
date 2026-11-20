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

// AddCephPrimaryStorage adds CephPrimaryStorage
func (cli *ZSClient) AddCephPrimaryStorage(params param.AddCephPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post("v1/primary-storage/ceph", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryCephPrimaryStorage queries CephPrimaryStorage list
func (cli *ZSClient) QueryCephPrimaryStorage(params *param.QueryParam) ([]view.PrimaryStorageInventoryView, error) {
	var resp []view.PrimaryStorageInventoryView
	return resp, cli.List("v1/primary-storage/ceph", params, &resp)
}

func (cli *ZSClient) GetCephPrimaryStorage(uuid string) (*view.PrimaryStorageInventoryView, error) {
	var resp view.PrimaryStorageInventoryView
	if err := cli.Get("v1/primary-storage/ceph", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCephPrimaryStorage Pagination
func (cli *ZSClient) PageCephPrimaryStorage(params *param.QueryParam) ([]view.PrimaryStorageInventoryView, int, error) {
	var cephPrimaryStorages []view.PrimaryStorageInventoryView
	total, err := cli.Page("v1/primary-storage/ceph", params, &cephPrimaryStorages)
	return cephPrimaryStorages, total, err
}

// RemoveMonFromCephPrimaryStorage requires additional arguments.
//
// Deprecated: use RemoveMonFromCephPrimaryStorageWithParams to supply the required fields.
func (cli *ZSClient) RemoveMonFromCephPrimaryStorage(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveMonFromCephPrimaryStorage requires additional arguments; use RemoveMonFromCephPrimaryStorageWithParams")
}

// RemoveMonFromCephPrimaryStorageWithParams sends the required route and query parameters for RemoveMonFromCephPrimaryStorage.
func (cli *ZSClient) RemoveMonFromCephPrimaryStorageWithParams(uuid string, deleteMode param.DeleteMode, params param.RemoveMonFromCephPrimaryStorageParamDetail) error {
	query := url.Values{
		"monHostnames": params.MonHostnames,
	}
	return cli.deleteActionWithQuery("v1/primary-storage/ceph", uuid, "mons", deleteMode, query, "monHostnames")
}

// AddMonToCephPrimaryStorage adds MonToCephPrimaryStorage
func (cli *ZSClient) AddMonToCephPrimaryStorage(uuid string, params param.AddMonToCephPrimaryStorageParam) (*view.CephPrimaryStorageInventoryView, error) {
	resp := view.CephPrimaryStorageInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/primary-storage/ceph/%s/mons", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
