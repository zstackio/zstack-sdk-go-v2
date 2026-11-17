// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSNSTopic deletes SNSTopic
func (cli *ZSClient) DeleteSNSTopic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/topics", uuid, string(deleteMode))
}
// UpdateSNSTopic updates SNSTopic
func (cli *ZSClient) UpdateSNSTopic(uuid string, params param.UpdateSNSTopicParam) (*view.SNSTopicInventoryView, error) {
	resp := view.SNSTopicInventoryView{}
	if err := cli.PutWithRespKey("v1/sns/topics", uuid, "", map[string]interface{}{
		"updateSNSTopic": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSTopic queries SNSTopic list
func (cli *ZSClient) QuerySNSTopic(params *param.QueryParam) ([]view.SNSTopicInventoryView, error) {
	var resp []view.SNSTopicInventoryView
	return resp, cli.List("v1/sns/topics", params, &resp)
}

func (cli *ZSClient) GetSNSTopic(uuid string) (*view.SNSTopicInventoryView, error) {
	var resp view.SNSTopicInventoryView
	if err := cli.Get("v1/sns/topics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSTopic Pagination
func (cli *ZSClient) PageSNSTopic(params *param.QueryParam) ([]view.SNSTopicInventoryView, int, error) {
	var sNSTopics []view.SNSTopicInventoryView
	total, err := cli.Page("v1/sns/topics", params, &sNSTopics)
	return sNSTopics, total, err
}
// CreateSNSTopic creates SNSTopic
func (cli *ZSClient) CreateSNSTopic(params param.CreateSNSTopicParam) (*view.SNSTopicInventoryView, error) {
	resp := view.SNSTopicInventoryView{}
	if err := cli.Post("v1/sns/topics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsubscribeSNSTopic operates on UnsubscribeSNSTopic
func (cli *ZSClient) UnsubscribeSNSTopic(topicUuid string, endpointUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/sns/topics", topicUuid, fmt.Sprintf("endpoints/%s", endpointUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// SubscribeSNSTopic operates on SubscribeSNSTopic
func (cli *ZSClient) SubscribeSNSTopic(topicUuid string, endpointUuid string, params param.SubscribeSNSTopicParam) (*view.SubscribeSNSTopicEventView, error) {
	resp := view.SubscribeSNSTopicEventView{}
	if err := cli.Post(fmt.Sprintf("v1/sns/topics/%s/endpoints/%s", topicUuid, endpointUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSNSTopicState changes SNSTopicState
func (cli *ZSClient) ChangeSNSTopicState(uuid string, params param.ChangeSNSTopicStateParam) (*view.SNSTopicInventoryView, error) {
	resp := view.SNSTopicInventoryView{}
	if err := cli.PutWithRespKey("v1/zwatch/topics", uuid, "", map[string]interface{}{
		"changeSNSTopicState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QuerySNSTopicSubscriber queries SNSTopicSubscriber list
func (cli *ZSClient) QuerySNSTopicSubscriber(params *param.QueryParam) ([]view.SNSSubscriberInventoryView, error) {
	var resp []view.SNSSubscriberInventoryView
	return resp, cli.List("v1/sns/topics/subscribers", params, &resp)
}

func (cli *ZSClient) GetSNSTopicSubscriber(uuid string) (*view.SNSSubscriberInventoryView, error) {
	var resp view.SNSSubscriberInventoryView
	if err := cli.Get("v1/sns/topics/subscribers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSTopicSubscriber Pagination
func (cli *ZSClient) PageSNSTopicSubscriber(params *param.QueryParam) ([]view.SNSSubscriberInventoryView, int, error) {
	var sNSTopicSubscribers []view.SNSSubscriberInventoryView
	total, err := cli.Page("v1/sns/topics/subscribers", params, &sNSTopicSubscribers)
	return sNSTopicSubscribers, total, err
}
