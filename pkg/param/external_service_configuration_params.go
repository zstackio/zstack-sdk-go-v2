// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// UpdateExternalServiceConfigurationParamDetail UpdateExternalServiceConfiguration detail param
type UpdateExternalServiceConfigurationParamDetail struct {
	Description *string `json:"description" validate:"required"`
}

// UpdateExternalServiceConfigurationParam UpdateExternalServiceConfiguration request param
type UpdateExternalServiceConfigurationParam struct {
	BaseParam
	Params UpdateExternalServiceConfigurationParamDetail `json:"updateExternalServiceConfiguration"`
}
// AddExternalServiceConfigurationParamDetail AddExternalServiceConfiguration detail param
type AddExternalServiceConfigurationParamDetail struct {
	ExternalServiceType string `json:"externalServiceType" validate:"required"`
	Configuration string `json:"configuration" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddExternalServiceConfigurationParam AddExternalServiceConfiguration request param
type AddExternalServiceConfigurationParam struct {
	BaseParam
	Params AddExternalServiceConfigurationParamDetail `json:"params"`
}
// DeleteExternalServiceConfigurationParamDetail DeleteExternalServiceConfiguration detail param
type DeleteExternalServiceConfigurationParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteExternalServiceConfigurationParam DeleteExternalServiceConfiguration request param
type DeleteExternalServiceConfigurationParam struct {
	BaseParam
	Params DeleteExternalServiceConfigurationParamDetail `json:"deleteExternalServiceConfiguration"`
}
