// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSEmaySmsEndpointInventoryView SNSEmaySmsEndpoint
type SNSEmaySmsEndpointInventoryView struct {
	BaseInfoView
	BaseTimeView
	RequestUrl string `json:"requestUrl,omitempty"`
	SmsAccessKeyId string `json:"smsAccessKeyId,omitempty"`
	SmsAccessKeySecret string `json:"smsAccessKeySecret,omitempty"`
	Supplier string `json:"supplier,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Platform SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

