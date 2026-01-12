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

func (cli *ZSClient) GetCdpTask(uuid string) (*view.CdpTaskInventoryView, error) {
	var resp view.CdpTaskInventoryView
	if err := cli.Get("v1/cdp-task", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateCdpTask updates CdpTask
func (cli *ZSClient) UpdateCdpTask(uuid string, params param.UpdateCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	var resp view.UpdateCdpTaskEventView
	err := cli.PutWithSpec("v1/cdp-backup-storage/task", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteCdpTask deletes CdpTask
func (cli *ZSClient) DeleteCdpTask(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/cdp-task", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
