// Copyright (c) ZStack.io, Inc.

package view

import "time"

// SNSEndpointThirdpartyAlertHistoryInventoryView SNSEndpointThirdpartyAlertHistory
type SNSEndpointThirdpartyAlertHistoryInventoryView struct {
	rest string `json:"alertUuid,omitempty"`
	rest string `json:"endpointUuid,omitempty"`
	rest string `json:"subscriptionUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
}

