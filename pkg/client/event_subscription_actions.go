// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryEventSubscription queries EventSubscription list
func (cli *ZSClient) QueryEventSubscription(params *param.QueryParam) ([]view.EventSubscriptionInventoryView, error) {
	var resp []view.EventSubscriptionInventoryView
	return resp, cli.List("v1/zwatch/events/subscriptions", params, &resp)
}

func (cli *ZSClient) GetEventSubscription(uuid string) (*view.EventSubscriptionInventoryView, error) {
	var resp view.EventSubscriptionInventoryView
	if err := cli.Get("v1/zwatch/events/subscriptions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEventSubscription Pagination
func (cli *ZSClient) PageEventSubscription(params *param.QueryParam) ([]view.EventSubscriptionInventoryView, int, error) {
	var eventSubscriptions []view.EventSubscriptionInventoryView
	total, err := cli.Page("v1/zwatch/events/subscriptions", params, &eventSubscriptions)
	return eventSubscriptions, total, err
}

// SubscribeEvent operates on SubscribeEvent
func (cli *ZSClient) SubscribeEvent(params param.SubscribeEventParam) (*view.EventSubscriptionInventoryView, error) {
	resp := view.EventSubscriptionInventoryView{}
	if err := cli.Post("v1/zwatch/events/subscriptions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsubscribeEvent operates on UnsubscribeEvent
func (cli *ZSClient) UnsubscribeEvent(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/events/subscriptions", uuid, string(deleteMode))
}

// UpdateEventData updates EventData
func (cli *ZSClient) UpdateEventData(params param.UpdateEventDataParam) (*view.UpdateEventDataEventView, error) {
	resp := view.UpdateEventDataEventView{}
	if err := cli.PutWithRespKey("v1/zwatch/events/actions", "", "", map[string]interface{}{
		"updateEventData": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAllEventMetadata gets AllEventMetadata by uuid
func (cli *ZSClient) GetAllEventMetadata() (*view.GetAllEventMetadataView, error) {
	var resp view.GetAllEventMetadataView
	if err := cli.GetWithRespKey("v1/zwatch/events/meta-data", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeEventSubscriptionState changes EventSubscriptionState
func (cli *ZSClient) ChangeEventSubscriptionState(uuid string, params param.ChangeEventSubscriptionStateParam) (*view.EventSubscriptionInventoryView, error) {
	resp := view.EventSubscriptionInventoryView{}
	if err := cli.PutWithRespKey("v1/zwatch/change/eventSubscription", uuid, "", map[string]interface{}{
		"changeEventSubscriptionState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveActionFromEventSubscription removes ActionFromEventSubscription
func (cli *ZSClient) RemoveActionFromEventSubscription(subscriptionUuid string, actionUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/zwatch/events/subscriptions", subscriptionUuid, fmt.Sprintf("actions/%s", actionUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AddActionToEventSubscription adds ActionToEventSubscription
func (cli *ZSClient) AddActionToEventSubscription(subscriptionUuid string, params param.AddActionToEventSubscriptionParam) (*view.EventSubscriptionInventoryView, error) {
	resp := view.EventSubscriptionInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/zwatch/events/subscriptions/%s/actions", subscriptionUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEventData gets EventData by uuid
func (cli *ZSClient) GetEventData() (*view.GetEventDataView, error) {
	var resp view.GetEventDataView
	if err := cli.GetWithRespKey("v1/zwatch/events", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSubscribeEvent updates SubscribeEvent
func (cli *ZSClient) UpdateSubscribeEvent(uuid string, params param.UpdateSubscribeEventParam) (*view.EventSubscriptionInventoryView, error) {
	resp := view.EventSubscriptionInventoryView{}
	if err := cli.PutWithRespKey("v1/zwatch/events/subscriptions", uuid, "", map[string]interface{}{
		"updateSubscribeEvent": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveLabelFromEventSubscription removes LabelFromEventSubscription
func (cli *ZSClient) RemoveLabelFromEventSubscription(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/events/subscriptions/labels", uuid, string(deleteMode))
}
