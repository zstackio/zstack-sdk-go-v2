// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSUniversalSmsEndpointInventoryView SNSUniversalSmsEndpoint
type SNSUniversalSmsEndpointInventoryView struct {
	SmsAccessKeyId string `json:"smsAccessKeyId,omitempty"`
	SmsAccessKeySecret string `json:"smsAccessKeySecret,omitempty"`
	Supplier string `json:"supplier,omitempty"`
	Name string `json:"name,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	PlatformUuid string `json:"platformUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	ConnectionStatus string `json:"connectionStatus,omitempty"`
	Platform SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

