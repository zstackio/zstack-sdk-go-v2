// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CleanLongJob operates on LongJob
func (cli *ZSClient) CleanLongJob(uuid string, params param.CleanLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	err := cli.PutWithSpec("v1/longjobs", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
// ResumeLongJob operates on LongJob
func (cli *ZSClient) ResumeLongJob(uuid string, params param.ResumeLongJobParam) (*view.LongJobInventoryView, error) {
	var resp view.ResumeLongJobEventView
	err := cli.PutWithSpec("v1/longjobs", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteLongJob deletes LongJob
func (cli *ZSClient) DeleteLongJob(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/longjobs", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// UpdateLongJob updates LongJob
func (cli *ZSClient) UpdateLongJob(uuid string, params param.UpdateLongJobParam) (*view.LongJobInventoryView, error) {
	var resp view.UpdateLongJobEventView
	err := cli.PutWithSpec("v1/longjobs", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryLongJob queries LongJob list
func (cli *ZSClient) QueryLongJob(params *param.QueryParam) ([]view.LongJobInventoryView, error) {
	var resp []view.LongJobInventoryView
	return resp, cli.List("v1/longjobs", params, &resp)
}

func (cli *ZSClient) GetLongJob(uuid string) (*view.LongJobInventoryView, error) {
	var resp view.LongJobInventoryView
	if err := cli.Get("v1/longjobs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
