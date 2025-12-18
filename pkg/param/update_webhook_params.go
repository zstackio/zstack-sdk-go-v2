// Copyright (c) ZStack.io, Inc.

package param

// UpdateWebhookDetailParam UpdateWebhook detail param
type UpdateWebhookDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty"`
	Type string `json:"type,omitempty"`
	Opaque string `json:"opaque,omitempty"`
}

// UpdateWebhookParam UpdateWebhook request param
type UpdateWebhookParam struct {
	BaseParam
	Params UpdateWebhookDetailParam `json:"params"`
}
