// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateEventSubscriptionLabel 更新EventSubscriptionLabel
func (cli *ZSClient) UpdateEventSubscriptionLabel(uuid string, params param.UpdateEventSubscriptionLabelParam) (*view.UpdateEventSubscriptionLabelEventView, error) {
	resp := view.UpdateEventSubscriptionLabelEventView{}
	if err := cli.Put("v1/zwatch/events/subscriptions/labels/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

