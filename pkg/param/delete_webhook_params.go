// Copyright (c) ZStack.io, Inc.

package param

// DeleteWebhookDetailParam DeleteWebhook detail param
type DeleteWebhookDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteWebhookParam DeleteWebhook request param
type DeleteWebhookParam struct {
	BaseParam
	Params DeleteWebhookDetailParam `json:"params"`
}
