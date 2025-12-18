// Copyright (c) ZStack.io, Inc.

package param

// GetObservabilityServerServiceDataDetailParam GetObservabilityServerServiceData详细参数
type GetObservabilityServerServiceDataDetailParam struct {
	rest string `json:"observabilityServerUuid" validate:"required"` // 必填
	rest string `json:"serviceType" validate:"required"` // 必填
	rest string `json:"serviceUuid" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest string `json:"startTime,omitempty"`
	rest string `json:"endTime,omitempty"`
	rest string `json:"sortDirection,omitempty"`
	rest map[string]string `json:"labelFilters,omitempty"`
}

// GetObservabilityServerServiceDataParam GetObservabilityServerServiceData请求参数
type GetObservabilityServerServiceDataParam struct {
	BaseParam
	Params GetObservabilityServerServiceDataDetailParam `json:"params"` // 详细参数
}

