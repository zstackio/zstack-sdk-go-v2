// Copyright (c) ZStack.io, Inc.

package param

// CreateWebhookDetailParam CreateWebhook详细参数
type CreateWebhookDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"url" validate:"required"` // 必填
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"opaque,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateWebhookParam CreateWebhook请求参数
type CreateWebhookParam struct {
	BaseParam
	Params CreateWebhookDetailParam `json:"params"` // 详细参数
}

