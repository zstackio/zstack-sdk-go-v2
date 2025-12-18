// Copyright (c) ZStack.io, Inc.

package param

// PrometheusQueryVmMonitoringDataDetailParam PrometheusQueryVmMonitoringData详细参数
type PrometheusQueryVmMonitoringDataDetailParam struct {
	rest []string `json:"vmUuids" validate:"required"` // 必填
	rest bool `json:"instant,omitempty"`
	rest int64 `json:"startTime,omitempty"`
	rest int64 `json:"endTime,omitempty"`
	rest string `json:"step,omitempty"`
	rest string `json:"expression" validate:"required"` // 必填
	rest string `json:"relativeTime,omitempty"`
}

// PrometheusQueryVmMonitoringDataParam PrometheusQueryVmMonitoringData请求参数
type PrometheusQueryVmMonitoringDataParam struct {
	BaseParam
	Params PrometheusQueryVmMonitoringDataDetailParam `json:"params"` // 详细参数
}

