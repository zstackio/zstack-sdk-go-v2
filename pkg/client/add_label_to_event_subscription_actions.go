// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddLabelToEventSubscription 操作AddLabelToEventSubscription
func (cli *ZSClient) AddLabelToEventSubscription(params param.AddLabelToEventSubscriptionParam) (*view.AddLabelToEventSubscriptionEventView, error) {
	resp := view.AddLabelToEventSubscriptionEventView{}
	if err := cli.Post("v1/zwatch/events/subscriptions/{subscriptionUuid}/labels", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

