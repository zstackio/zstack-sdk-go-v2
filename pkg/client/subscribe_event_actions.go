// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SubscribeEvent operates on SubscribeEvent
func (cli *ZSClient) SubscribeEvent(params param.SubscribeEventParam) (*view.SubscribeEventEventView, error) {
	resp := view.SubscribeEventEventView{}
	if err := cli.Post("v1/zwatch/events/subscriptions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
