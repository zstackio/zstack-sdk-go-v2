// Copyright (c) ZStack.io, Inc.

package view

// SNSEmailTestConnectionEventView SNSEmailTestConnectionEvent
type SNSEmailTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

