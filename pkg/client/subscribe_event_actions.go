// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SubscribeEvent operates on SubscribeEvent
func (cli *ZSClient) SubscribeEvent(params param.SubscribeEventParam) (*view.SubscribeEventEventView, error) {
	resp := view.SubscribeEventEventView{}
	if err := cli.Post("v1/zwatch/events/subscriptions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
