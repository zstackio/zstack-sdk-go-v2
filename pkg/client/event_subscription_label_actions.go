// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateEventSubscriptionLabel updates EventSubscriptionLabel
func (cli *ZSClient) UpdateEventSubscriptionLabel(uuid string, params param.UpdateEventSubscriptionLabelParam) (*view.EventSubscriptionLabelInventoryView, error) {
	resp := view.EventSubscriptionLabelInventoryView{}
	if err := cli.PutWithRespKey("v1/zwatch/events/subscriptions/labels", uuid, "", map[string]interface{}{
		"updateEventSubscriptionLabel": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddLabelToEventSubscription adds LabelToEventSubscription
func (cli *ZSClient) AddLabelToEventSubscription(subscriptionUuid string, params param.AddLabelToEventSubscriptionParam) (*view.EventSubscriptionLabelInventoryView, error) {
	resp := view.EventSubscriptionLabelInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/zwatch/events/subscriptions/%s/labels", subscriptionUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
