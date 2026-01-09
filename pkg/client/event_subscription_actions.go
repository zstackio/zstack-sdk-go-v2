// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryEventSubscription queries EventSubscription list
func (cli *ZSClient) QueryEventSubscription(params *param.QueryParam) ([]view.EventSubscriptionInventoryView, error) {
	var resp []view.EventSubscriptionInventoryView
	return resp, cli.List("v1/zwatch/events/subscriptions", params, &resp)
}

func (cli *ZSClient) GetEventSubscription(uuid string) (*view.EventSubscriptionInventoryView, error) {
	var resp view.EventSubscriptionInventoryView
	if err := cli.Get("v1/zwatch/events/subscriptions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
