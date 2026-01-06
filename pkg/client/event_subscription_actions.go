// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryEventSubscription queries EventSubscription list
func (cli *ZSClient) QueryEventSubscription(params *param.QueryParam) ([]view.EventSubscriptionInventoryView, error) {
	var resp []view.EventSubscriptionInventoryView
	return resp, cli.List("v1/zwatch/events/subscriptions", params, &resp)
}
