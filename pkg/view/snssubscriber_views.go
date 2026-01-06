// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSSubscriberInventoryView SNSSubscriber
type SNSSubscriberInventoryView struct {
	TopicUuid string `json:"topicUuid,omitempty"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// QuerySNSTopicSubscriberView QuerySNSTopicSubscriber
type QuerySNSTopicSubscriberView struct {
	Inventories []SNSSubscriberInventoryView `json:"inventories,omitempty"`
}

