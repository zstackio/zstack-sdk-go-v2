// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
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
	return cli.Delete("v1/cbt-task/{uuid}", uuid, string(deleteMode))
}
// QueryCbtTask queries CbtTask list
func (cli *ZSClient) QueryCbtTask(params *param.QueryParam) ([]view.CbtTaskInventoryView, error) {
	var resp []view.CbtTaskInventoryView
	return resp, cli.List("v1/cbt-task", params, &resp)
}
