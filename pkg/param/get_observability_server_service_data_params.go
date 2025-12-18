// Copyright (c) ZStack.io, Inc.

package param

// GetObservabilityServerServiceDataDetailParam GetObservabilityServerServiceData detail param
type GetObservabilityServerServiceDataDetailParam struct {
	ObservabilityServerUuid string `json:"observabilityServerUuid" validate:"required"`
	ServiceType string `json:"serviceType" validate:"required"`
	ServiceUuid string `json:"serviceUuid" validate:"required"`
	Limit int `json:"limit,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	EndTime string `json:"endTime,omitempty"`
	SortDirection string `json:"sortDirection,omitempty"`
	LabelFilters map[string]string `json:"labelFilters,omitempty"`
}

// GetObservabilityServerServiceDataParam GetObservabilityServerServiceData request param
type GetObservabilityServerServiceDataParam struct {
	BaseParam
	Params GetObservabilityServerServiceDataDetailParam `json:"params"`
}
