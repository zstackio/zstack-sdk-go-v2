// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPublishApp queries PublishApp list
func (cli *ZSClient) QueryPublishApp(ctx context.Context, params *param.QueryParam) ([]view.PublishAppInventoryView, error) {
	var resp []view.PublishAppInventoryView
	return resp, cli.List(ctx, "v1/appcenter/app", params, &resp)
}

func (cli *ZSClient) GetPublishApp(ctx context.Context, uuid string) (*view.PublishAppInventoryView, error) {
	var resp view.PublishAppInventoryView
	if err := cli.Get(ctx, "v1/appcenter/app", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePublishApp Pagination
func (cli *ZSClient) PagePublishApp(ctx context.Context, params *param.QueryParam) ([]view.PublishAppInventoryView, int, error) {
	var publishApps []view.PublishAppInventoryView
	total, err := cli.Page(ctx, "v1/appcenter/app", params, &publishApps)
	return publishApps, total, err
}
// PublishApp operates on PublishApp
func (cli *ZSClient) PublishApp(ctx context.Context, params param.PublishAppParam) (*view.PublishAppInventoryView, error) {
	resp := view.PublishAppInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/appcenter/app", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdatePublishApp updates PublishApp
func (cli *ZSClient) UpdatePublishApp(ctx context.Context, uuid string, params param.UpdatePublishAppParam) (*view.PublishAppInventoryView, error) {
	resp := view.PublishAppInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/appcenter/app", uuid, "", map[string]interface{}{
		"updatePublishApp": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePublishApp deletes PublishApp
func (cli *ZSClient) DeletePublishApp(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/appcenter/app", uuid, string(deleteMode))
}
