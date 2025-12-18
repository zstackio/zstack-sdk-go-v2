// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SNSSubscriberInventoryView SNSSubscriber
type SNSSubscriberInventoryView struct {
	rest string `json:"topicUuid,omitempty"`
	rest string `json:"endpointUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

