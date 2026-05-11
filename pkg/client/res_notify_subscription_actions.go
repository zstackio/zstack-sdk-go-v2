// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteResNotifySubscription deletes ResNotifySubscription
func (cli *ZSClient) DeleteResNotifySubscription(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/resnotify/subscriptions", uuid, string(deleteMode))
}
// QueryResNotifySubscription queries ResNotifySubscription list
func (cli *ZSClient) QueryResNotifySubscription(ctx context.Context, params *param.QueryParam) ([]view.ResNotifySubscriptionInventoryView, error) {
	var resp []view.ResNotifySubscriptionInventoryView
	return resp, cli.List(ctx, "v1/zwatch/resnotify/subscriptions", params, &resp)
}

func (cli *ZSClient) GetResNotifySubscription(ctx context.Context, uuid string) (*view.ResNotifySubscriptionInventoryView, error) {
	var resp view.ResNotifySubscriptionInventoryView
	if err := cli.Get(ctx, "v1/zwatch/resnotify/subscriptions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageResNotifySubscription Pagination
func (cli *ZSClient) PageResNotifySubscription(ctx context.Context, params *param.QueryParam) ([]view.ResNotifySubscriptionInventoryView, int, error) {
	var resNotifySubscriptions []view.ResNotifySubscriptionInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/resnotify/subscriptions", params, &resNotifySubscriptions)
	return resNotifySubscriptions, total, err
}
// UpdateResNotifySubscription updates ResNotifySubscription
func (cli *ZSClient) UpdateResNotifySubscription(ctx context.Context, uuid string, params param.UpdateResNotifySubscriptionParam) (*view.ResNotifySubscriptionInventoryView, error) {
	resp := view.ResNotifySubscriptionInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/zwatch/resnotify/subscriptions", uuid, "", map[string]interface{}{
		"updateResNotifySubscription": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
