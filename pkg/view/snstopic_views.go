// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSTopicInventoryView SNSTopic
type SNSTopicInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Locale string `json:"locale,omitempty"`
	CreateDate ZStackTime `json:"createDate,omitempty"`
	LastOpDate ZStackTime `json:"lastOpDate,omitempty"`
}

// DeleteSNSTopicEventView DeleteSNSTopicEvent
type DeleteSNSTopicEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateSNSTopicEventView UpdateSNSTopicEvent
type UpdateSNSTopicEventView struct {
	Inventory SNSTopicInventoryView `json:"inventory,omitempty"`
}

// QuerySNSTopicView QuerySNSTopic
type QuerySNSTopicView struct {
	Inventories []SNSTopicInventoryView `json:"inventories,omitempty"`
}

// ChangeSNSTopicStateEventView ChangeSNSTopicStateEvent
type ChangeSNSTopicStateEventView struct {
	Inventory SNSTopicInventoryView `json:"inventory,omitempty"`
}

// CreateSNSTopicEventView CreateSNSTopicEvent
type CreateSNSTopicEventView struct {
	Inventory SNSTopicInventoryView `json:"inventory,omitempty"`
}

