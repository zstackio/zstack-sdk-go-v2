// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSubscribeEvent updates SubscribeEvent
func (cli *ZSClient) UpdateSubscribeEvent(uuid string, params param.UpdateSubscribeEventParam) (*view.UpdateSubscribeEventEventView, error) {
	resp := view.UpdateSubscribeEventEventView{}
	if err := cli.Put("v1/zwatch/events/subscriptions/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
