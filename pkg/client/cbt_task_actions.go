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
	resp := view.CbtTaskInventoryView{}
	if err := cli.Post("v1/cbt-task/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteCbtTask deletes CbtTask
func (cli *ZSClient) DeleteCbtTask(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cbt-task", uuid, string(deleteMode))
}
// QueryCbtTask queries CbtTask list
func (cli *ZSClient) QueryCbtTask(params *param.QueryParam) ([]view.CbtTaskInventoryView, error) {
	var resp []view.CbtTaskInventoryView
	return resp, cli.List("v1/cbt-task", params, &resp)
}

// PageCbtTask Pagination
func (cli *ZSClient) PageCbtTask(params *param.QueryParam) ([]view.CbtTaskInventoryView, int, error) {
	var cbtTasks []view.CbtTaskInventoryView
	total, err := cli.Page("v1/cbt-task", params, &cbtTasks)
	return cbtTasks, total, err
}
