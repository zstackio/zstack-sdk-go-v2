// Copyright (c) ZStack.io, Inc.

package view

// SNSMicrosoftTeamsTestConnectionEventView SNSMicrosoftTeamsTestConnectionEvent
type SNSMicrosoftTeamsTestConnectionEventView struct {
	Connected bool `json:"connected,omitempty"`
	WebhookResp map[string]interface{} `json:"webhookResp,omitempty"`
}

