// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// SetupExternalAlertIntegrationParamDetail SetupExternalAlertIntegration detail param
type SetupExternalAlertIntegrationParamDetail struct {
	Name        string  `json:"name"`                  // 集成名称
	Url         string  `json:"url"`                   // Webhook 回调地址
	Description *string `json:"description,omitempty"` // 描述
}

// SetupExternalAlertIntegrationParam SetupExternalAlertIntegration request param
type SetupExternalAlertIntegrationParam struct {
	BaseParam
	Params SetupExternalAlertIntegrationParamDetail `json:"params"`
}

// CheckExternalAlertIntegrationParamDetail CheckExternalAlertIntegration detail param
type CheckExternalAlertIntegrationParamDetail struct {
	EndpointUuid string `json:"endpointUuid"` // HTTP Endpoint UUID
	ExpectedUrl  string `json:"expectedUrl"`  // 期望的 Webhook 回调地址
}

// CheckExternalAlertIntegrationParam CheckExternalAlertIntegration request param
type CheckExternalAlertIntegrationParam struct {
	BaseParam
	Params CheckExternalAlertIntegrationParamDetail `json:"params"`
}

// RemoveExternalAlertIntegrationParamDetail RemoveExternalAlertIntegration detail param
type RemoveExternalAlertIntegrationParamDetail struct {
	EndpointUuid string `json:"endpointUuid"` // HTTP Endpoint UUID
}

// RemoveExternalAlertIntegrationParam RemoveExternalAlertIntegration request param
type RemoveExternalAlertIntegrationParam struct {
	BaseParam
	Params RemoveExternalAlertIntegrationParamDetail `json:"params"`
}
