// Copyright (c) ZStack.io, Inc.

package param

// CreateMetricTemplateDetailParam CreateMetricTemplate detail param
type CreateMetricTemplateDetailParam struct {
	ReceiverUuid string `json:"receiverUuid" validate:"required"`
	Template string `json:"template" validate:"required"`
	Namespace string `json:"namespace" validate:"required"`
	MetricName string `json:"metricName" validate:"required"`
	LabelsJsonStr string `json:"labelsJsonStr,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateMetricTemplateParam CreateMetricTemplate request param
type CreateMetricTemplateParam struct {
	BaseParam
	Params CreateMetricTemplateDetailParam `json:"params"`
}
