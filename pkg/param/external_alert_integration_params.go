// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// SetupExternalAlertIntegrationDetail SetupExternalAlertIntegration detail param
type SetupExternalAlertIntegrationDetail struct {
	Name        string  `json:"name"`                  // 集成名称
	Url         string  `json:"url"`                   // Webhook 回调地址
	Description *string `json:"description,omitempty"` // 描述
}

// SetupExternalAlertIntegrationParam SetupExternalAlertIntegration request param
type SetupExternalAlertIntegrationParam struct {
	BaseParam
	Params SetupExternalAlertIntegrationDetail `json:"params"`
}

// CheckExternalAlertIntegrationDetail CheckExternalAlertIntegration detail param
type CheckExternalAlertIntegrationDetail struct {
	EndpointUuid string `json:"endpointUuid"` // HTTP Endpoint UUID
	ExpectedUrl  string `json:"expectedUrl"`  // 期望的 Webhook 回调地址
}

// CheckExternalAlertIntegrationParam CheckExternalAlertIntegration request param
type CheckExternalAlertIntegrationParam struct {
	BaseParam
	Params CheckExternalAlertIntegrationDetail `json:"params"`
}

// RemoveExternalAlertIntegrationDetail RemoveExternalAlertIntegration detail param
type RemoveExternalAlertIntegrationDetail struct {
	EndpointUuid string `json:"endpointUuid"` // HTTP Endpoint UUID
}

// RemoveExternalAlertIntegrationParam RemoveExternalAlertIntegration request param
type RemoveExternalAlertIntegrationParam struct {
	BaseParam
	Params RemoveExternalAlertIntegrationDetail `json:"params"`
}
