// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateCdpTask creates CdpTask
func (cli *ZSClient) CreateCdpTask(params param.CreateCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	resp := view.CdpTaskInventoryView{}
	if err := cli.Post("v1/cdp-backup-storage/task", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryCdpTask queries CdpTask list
func (cli *ZSClient) QueryCdpTask(params *param.QueryParam) ([]view.CdpTaskInventoryView, error) {
	var resp []view.CdpTaskInventoryView
	return resp, cli.List("v1/cdp-task", params, &resp)
}

func (cli *ZSClient) GetCdpTask(uuid string) (*view.CdpTaskInventoryView, error) {
	var resp view.CdpTaskInventoryView
	if err := cli.Get("v1/cdp-task", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCdpTask Pagination
func (cli *ZSClient) PageCdpTask(params *param.QueryParam) ([]view.CdpTaskInventoryView, int, error) {
	var cdpTasks []view.CdpTaskInventoryView
	total, err := cli.Page("v1/cdp-task", params, &cdpTasks)
	return cdpTasks, total, err
}
// UpdateCdpTask updates CdpTask
func (cli *ZSClient) UpdateCdpTask(uuid string, params param.UpdateCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	resp := view.CdpTaskInventoryView{}
	if err := cli.Put("v1/cdp-backup-storage/task", uuid, map[string]interface{}{
		"updateCdpTask": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteCdpTask deletes CdpTask
func (cli *ZSClient) DeleteCdpTask(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cdp-task", uuid, string(deleteMode))
}
