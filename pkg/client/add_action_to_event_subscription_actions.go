// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddActionToEventSubscription adds ActionToEventSubscription
func (cli *ZSClient) AddActionToEventSubscription(params param.AddActionToEventSubscriptionParam) (*view.AddActionToEventSubscriptionEventView, error) {
	resp := view.AddActionToEventSubscriptionEventView{}
	if err := cli.Post("v1/zwatch/events/subscriptions/{subscriptionUuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
