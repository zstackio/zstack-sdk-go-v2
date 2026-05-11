// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ResNotifyWebhookRefInventoryView ResNotifyWebhookRef
type ResNotifyWebhookRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	WebhookUrl string `json:"webhookUrl,omitempty"`
	CustomHeaders string `json:"customHeaders,omitempty"`
}

