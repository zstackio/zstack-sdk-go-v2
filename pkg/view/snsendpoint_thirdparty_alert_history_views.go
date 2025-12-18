// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSEndpointThirdpartyAlertHistoryInventoryView SNSEndpointThirdpartyAlertHistory
type SNSEndpointThirdpartyAlertHistoryInventoryView struct {
	AlertUuid string `json:"alertUuid,omitempty"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
	SubscriptionUuid string `json:"subscriptionUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
}

