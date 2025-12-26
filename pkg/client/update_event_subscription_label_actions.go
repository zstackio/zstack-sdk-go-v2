// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateEventSubscriptionLabel updates EventSubscriptionLabel
func (cli *ZSClient) UpdateEventSubscriptionLabel(uuid string, params param.UpdateEventSubscriptionLabelParam) (*view.UpdateEventSubscriptionLabelEventView, error) {
	resp := view.UpdateEventSubscriptionLabelEventView{}
	if err := cli.Put("v1/zwatch/events/subscriptions/labels/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
