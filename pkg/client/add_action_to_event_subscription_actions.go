// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddActionToEventSubscription adds ActionToEventSubscription
func (cli *ZSClient) AddActionToEventSubscription(params param.AddActionToEventSubscriptionParam) (*view.AddActionToEventSubscriptionEventView, error) {
	resp := view.AddActionToEventSubscriptionEventView{}
	if err := cli.Post("v1/zwatch/events/subscriptions/{subscriptionUuid}/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
