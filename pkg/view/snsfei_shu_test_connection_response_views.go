// Copyright (c) ZStack.io, Inc.

package view

// SNSFeiShuTestConnectionEventView SNSFeiShuTestConnectionEvent
type SNSFeiShuTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

