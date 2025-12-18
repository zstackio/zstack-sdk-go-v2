// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeEventSubscriptionState 操作EventSubscriptionState
func (cli *ZSClient) ChangeEventSubscriptionState(uuid string, params param.ChangeEventSubscriptionStateParam) (*view.ChangeEventSubscriptionStateEventView, error) {
	resp := view.ChangeEventSubscriptionStateEventView{}
	if err := cli.Put("v1/zwatch/change/eventSubscription/{uuid}/state", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

