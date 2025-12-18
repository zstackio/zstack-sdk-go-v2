// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSubscribeEvent 更新SubscribeEvent
func (cli *ZSClient) UpdateSubscribeEvent(uuid string, params param.UpdateSubscribeEventParam) (*view.UpdateSubscribeEventEventView, error) {
	resp := view.UpdateSubscribeEventEventView{}
	if err := cli.Put("v1/zwatch/events/subscriptions/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

