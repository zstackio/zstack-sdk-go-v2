// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CleanLongJob operates on LongJob
func (cli *ZSClient) CleanLongJob(uuid string, params param.CleanLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ResumeLongJob operates on LongJob
func (cli *ZSClient) ResumeLongJob(uuid string, params param.ResumeLongJobParam) (*view.LongJobInventoryView, error) {
	var resp view.ResumeLongJobEventView
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteLongJob deletes LongJob
func (cli *ZSClient) DeleteLongJob(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/longjobs/{uuid}", uuid, string(deleteMode))
}
// UpdateLongJob updates LongJob
func (cli *ZSClient) UpdateLongJob(uuid string, params param.UpdateLongJobParam) (*view.LongJobInventoryView, error) {
	var resp view.UpdateLongJobEventView
	if err := cli.Put("v1/longjobs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryLongJob queries LongJob list
func (cli *ZSClient) QueryLongJob(params *param.QueryParam) ([]view.LongJobInventoryView, error) {
	var resp []view.LongJobInventoryView
	return resp, cli.List("v1/longjobs", params, &resp)
}
