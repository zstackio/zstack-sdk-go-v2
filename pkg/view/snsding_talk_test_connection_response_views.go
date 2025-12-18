// Copyright (c) ZStack.io, Inc.

package view

// SNSDingTalkTestConnectionEventView SNSDingTalkTestConnectionEvent
type SNSDingTalkTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

