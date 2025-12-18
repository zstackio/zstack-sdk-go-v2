// Copyright (c) ZStack.io, Inc.

package param

// CreateWebhookDetailParam CreateWebhook detail param
type CreateWebhookDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Url string `json:"url" validate:"required"`
	Type string `json:"type" validate:"required"`
	Opaque string `json:"opaque,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateWebhookParam CreateWebhook request param
type CreateWebhookParam struct {
	BaseParam
	Params CreateWebhookDetailParam `json:"params"`
}
