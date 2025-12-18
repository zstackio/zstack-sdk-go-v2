// Copyright (c) ZStack.io, Inc.

package view

// SNSSnmpTestConnectionEventView SNSSnmpTestConnectionEvent
type SNSSnmpTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

