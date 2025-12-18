// Copyright (c) ZStack.io, Inc.

package param

// PrometheusQueryVmMonitoringDataDetailParam PrometheusQueryVmMonitoringData detail param
type PrometheusQueryVmMonitoringDataDetailParam struct {
	VmUuids []string `json:"vmUuids" validate:"required"`
	Instant bool `json:"instant,omitempty"`
	StartTime int64 `json:"startTime,omitempty"`
	EndTime int64 `json:"endTime,omitempty"`
	Step string `json:"step,omitempty"`
	Expression string `json:"expression" validate:"required"`
	RelativeTime string `json:"relativeTime,omitempty"`
}

// PrometheusQueryVmMonitoringDataParam PrometheusQueryVmMonitoringData request param
type PrometheusQueryVmMonitoringDataParam struct {
	BaseParam
	Params PrometheusQueryVmMonitoringDataDetailParam `json:"params"`
}
