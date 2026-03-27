// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CleanLongJob operates on LongJob
func (cli *ZSClient) CleanLongJob(ctx context.Context, uuid string, params param.CleanLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/longjobs", uuid, "", map[string]interface{}{
		"cleanLongJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ResumeLongJob operates on LongJob
func (cli *ZSClient) ResumeLongJob(ctx context.Context, uuid string, params param.ResumeLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/longjobs", uuid, "", map[string]interface{}{
		"resumeLongJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteLongJob deletes LongJob
func (cli *ZSClient) DeleteLongJob(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/longjobs", uuid, string(deleteMode))
}
// UpdateLongJob updates LongJob
func (cli *ZSClient) UpdateLongJob(ctx context.Context, uuid string, params param.UpdateLongJobParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/longjobs", uuid, "", map[string]interface{}{
		"updateLongJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryLongJob queries LongJob list
func (cli *ZSClient) QueryLongJob(ctx context.Context, params *param.QueryParam) ([]view.LongJobInventoryView, error) {
	var resp []view.LongJobInventoryView
	return resp, cli.List(ctx, "v1/longjobs", params, &resp)
}

func (cli *ZSClient) GetLongJob(ctx context.Context, uuid string) (*view.LongJobInventoryView, error) {
	var resp view.LongJobInventoryView
	if err := cli.Get(ctx, "v1/longjobs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLongJob Pagination
func (cli *ZSClient) PageLongJob(ctx context.Context, params *param.QueryParam) ([]view.LongJobInventoryView, int, error) {
	var longJobs []view.LongJobInventoryView
	total, err := cli.Page(ctx, "v1/longjobs", params, &longJobs)
	return longJobs, total, err
}
