// Copyright (c) ZStack.io, Inc.

package param

// CreateAlarmDetailParam CreateAlarm detail param
type CreateAlarmDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ComparisonOperator string `json:"comparisonOperator" validate:"required"`
	Period int `json:"period,omitempty"`
	Namespace string `json:"namespace" validate:"required"`
	MetricName string `json:"metricName" validate:"required"`
	Threshold float64 `json:"threshold" validate:"required"`
	RepeatInterval int `json:"repeatInterval,omitempty"`
	Labels []LabelParam `json:"labels,omitempty"`
	Actions []ActionParamParam `json:"actions,omitempty"`
	RepeatCount int `json:"repeatCount,omitempty"`
	Type string `json:"type,omitempty"`
	EnableRecovery bool `json:"enableRecovery,omitempty"`
	EmergencyLevel string `json:"emergencyLevel,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAlarmParam CreateAlarm request param
type CreateAlarmParam struct {
	BaseParam
	Params CreateAlarmDetailParam `json:"params"`
}
