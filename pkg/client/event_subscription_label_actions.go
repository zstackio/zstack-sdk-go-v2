// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateEventSubscriptionLabel updates EventSubscriptionLabel
func (cli *ZSClient) UpdateEventSubscriptionLabel(uuid string, params param.UpdateEventSubscriptionLabelParam) (*view.EventSubscriptionLabelInventoryView, error) {
	var resp view.UpdateEventSubscriptionLabelEventView
	if err := cli.Put("v1/zwatch/events/subscriptions/labels/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
