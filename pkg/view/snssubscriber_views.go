// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSSubscriberInventoryView SNSSubscriber
type SNSSubscriberInventoryView struct {
	TopicUuid    string    `json:"topicUuid,omitempty"`
	EndpointUuid string    `json:"endpointUuid,omitempty"`
	CreateDate   time.Time `json:"createDate,omitempty"`
	LastOpDate   time.Time `json:"lastOpDate,omitempty"`
}

// QuerySNSTopicSubscriberView QuerySNSTopicSubscriber
type QuerySNSTopicSubscriberView struct {
	Inventories []SNSSubscriberInventoryView `json:"inventories,omitempty"`
}
