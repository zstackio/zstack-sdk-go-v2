// Copyright (c) ZStack.io, Inc.

package view

// SNSWeComTestConnectionEventView SNSWeComTestConnectionEvent
type SNSWeComTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

