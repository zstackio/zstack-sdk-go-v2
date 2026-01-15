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
	if err := cli.Put("v1/longjobs", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ResumeLongJob operates on LongJob
func (cli *ZSClient) ResumeLongJob(uuid string, params param.ResumeLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.Put("v1/longjobs", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteLongJob deletes LongJob
func (cli *ZSClient) DeleteLongJob(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/longjobs", uuid, string(deleteMode))
}
// UpdateLongJob updates LongJob
func (cli *ZSClient) UpdateLongJob(uuid string, params param.UpdateLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.Put("v1/longjobs", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryLongJob queries LongJob list
func (cli *ZSClient) QueryLongJob(params *param.QueryParam) ([]view.LongJobInventoryView, error) {
	var resp []view.LongJobInventoryView
	return resp, cli.List("v1/longjobs", params, &resp)
}

// PageLongJob Pagination
func (cli *ZSClient) PageLongJob(params *param.QueryParam) ([]view.LongJobInventoryView, int, error) {
	var longJobs []view.LongJobInventoryView
	total, err := cli.Page("v1/longjobs", params, &longJobs)
	return longJobs, total, err
}
