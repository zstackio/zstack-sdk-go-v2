// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// SNSAliyunSmsEndpointInventoryView SNSAliyunSmsEndpoint
type SNSAliyunSmsEndpointInventoryView struct {
	BaseInfoView
	BaseTimeView
	Receivers          []SNSSmsReceiverInventoryView       `json:"receivers,omitempty"`
	SmsAccessKeyId     string                              `json:"smsAccessKeyId,omitempty"`
	SmsAccessKeySecret string                              `json:"smsAccessKeySecret,omitempty"`
	Supplier           string                              `json:"supplier,omitempty"`
	Type               string                              `json:"type,omitempty"`
	State              string                              `json:"state,omitempty"`
	PlatformUuid       string                              `json:"platformUuid,omitempty"`
	ConnectionStatus   string                              `json:"connectionStatus,omitempty"`
	Platform           SNSApplicationPlatformInventoryView `json:"platform,omitempty"`
}

// CreateSNSAliyunSmsEndpointEventView CreateSNSAliyunSmsEndpointEvent
type CreateSNSAliyunSmsEndpointEventView struct {
	Inventory SNSAliyunSmsEndpointInventoryView `json:"inventory,omitempty"`
}

// ValidateSNSAliyunSmsEndpointEventView ValidateSNSAliyunSmsEndpointEvent
type ValidateSNSAliyunSmsEndpointEventView struct {
	Success bool `json:"success,omitempty"`
}

// QuerySNSSmsEndpointView QuerySNSSmsEndpoint
type QuerySNSSmsEndpointView struct {
	Inventories []SNSAliyunSmsEndpointInventoryView `json:"inventories,omitempty"`
}
