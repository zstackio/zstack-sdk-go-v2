// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteResNotifySubscriptionParamDetail DeleteResNotifySubscription detail param
type DeleteResNotifySubscriptionParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteResNotifySubscriptionParam DeleteResNotifySubscription request param
type DeleteResNotifySubscriptionParam struct {
	BaseParam
	Params DeleteResNotifySubscriptionParamDetail `json:"deleteResNotifySubscription"`
}
// UpdateResNotifySubscriptionParamDetail UpdateResNotifySubscription detail param
type UpdateResNotifySubscriptionParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ResourceTypes []string `json:"resourceTypes,omitempty"`
	EventTypes []string `json:"eventTypes,omitempty"`
	State *string `json:"state,omitempty"`
	WebhookUrl *string `json:"webhookUrl,omitempty"`
	Secret *string `json:"secret,omitempty"`
	CustomHeaders *string `json:"customHeaders,omitempty"`
}

// UpdateResNotifySubscriptionParam UpdateResNotifySubscription request param
type UpdateResNotifySubscriptionParam struct {
	BaseParam
	Params UpdateResNotifySubscriptionParamDetail `json:"updateResNotifySubscription"`
}
