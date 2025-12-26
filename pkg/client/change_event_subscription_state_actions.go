// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeEventSubscriptionState changes EventSubscriptionState
func (cli *ZSClient) ChangeEventSubscriptionState(uuid string, params param.ChangeEventSubscriptionStateParam) (*view.ChangeEventSubscriptionStateEventView, error) {
	resp := view.ChangeEventSubscriptionStateEventView{}
	if err := cli.Put("v1/zwatch/change/eventSubscription/{uuid}/state", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
