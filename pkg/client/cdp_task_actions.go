// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateCdpTask creates CdpTask
func (cli *ZSClient) CreateCdpTask(params param.CreateCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	var resp view.CreateCdpTaskEventView
	if err := cli.Post("v1/cdp-backup-storage/task", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryCdpTask queries CdpTask list
func (cli *ZSClient) QueryCdpTask(params *param.QueryParam) ([]view.CdpTaskInventoryView, error) {
	var resp []view.CdpTaskInventoryView
	return resp, cli.List("v1/cdp-task", params, &resp)
}
// UpdateCdpTask updates CdpTask
func (cli *ZSClient) UpdateCdpTask(uuid string, params param.UpdateCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	var resp view.UpdateCdpTaskEventView
	if err := cli.Put("v1/cdp-backup-storage/task/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteCdpTask deletes CdpTask
func (cli *ZSClient) DeleteCdpTask(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cdp-task/{uuid}", uuid, string(deleteMode))
}
