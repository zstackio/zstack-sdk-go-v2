// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateCbtTask creates CbtTask
func (cli *ZSClient) CreateCbtTask(params param.CreateCbtTaskParam) (*view.CbtTaskInventoryView, error) {
	var resp view.CreateCbtTaskEventView
	if err := cli.Post("v1/cbt-task/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteCbtTask deletes CbtTask
func (cli *ZSClient) DeleteCbtTask(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/cbt-task", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// QueryCbtTask queries CbtTask list
func (cli *ZSClient) QueryCbtTask(params *param.QueryParam) ([]view.CbtTaskInventoryView, error) {
	var resp []view.CbtTaskInventoryView
	return resp, cli.List("v1/cbt-task", params, &resp)
}

func (cli *ZSClient) GetCbtTask(uuid string) (*view.CbtTaskInventoryView, error) {
	var resp view.CbtTaskInventoryView
	if err := cli.Get("v1/cbt-task", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
