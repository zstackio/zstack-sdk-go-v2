// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteResNotifySubscription deletes ResNotifySubscription
func (cli *ZSClient) DeleteResNotifySubscription(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/resnotify/subscriptions", uuid, string(deleteMode))
}
// QueryResNotifySubscription queries ResNotifySubscription list
func (cli *ZSClient) QueryResNotifySubscription(params *param.QueryParam) ([]view.ResNotifySubscriptionInventoryView, error) {
	var resp []view.ResNotifySubscriptionInventoryView
	return resp, cli.List("v1/zwatch/resnotify/subscriptions", params, &resp)
}

func (cli *ZSClient) GetResNotifySubscription(uuid string) (*view.ResNotifySubscriptionInventoryView, error) {
	var resp view.ResNotifySubscriptionInventoryView
	if err := cli.Get("v1/zwatch/resnotify/subscriptions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageResNotifySubscription Pagination
func (cli *ZSClient) PageResNotifySubscription(params *param.QueryParam) ([]view.ResNotifySubscriptionInventoryView, int, error) {
	var resNotifySubscriptions []view.ResNotifySubscriptionInventoryView
	total, err := cli.Page("v1/zwatch/resnotify/subscriptions", params, &resNotifySubscriptions)
	return resNotifySubscriptions, total, err
}
// UpdateResNotifySubscription updates ResNotifySubscription
func (cli *ZSClient) UpdateResNotifySubscription(uuid string, params param.UpdateResNotifySubscriptionParam) (*view.ResNotifySubscriptionInventoryView, error) {
	resp := view.ResNotifySubscriptionInventoryView{}
	if err := cli.PutWithRespKey("v1/zwatch/resnotify/subscriptions", uuid, "", map[string]interface{}{
		"updateResNotifySubscription": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
