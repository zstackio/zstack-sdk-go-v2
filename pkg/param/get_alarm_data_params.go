// Copyright (c) ZStack.io, Inc.

package param

// GetAlarmDataDetailParam GetAlarmData detail param
type GetAlarmDataDetailParam struct {
	StartTime int64 `json:"startTime,omitempty"`
	EndTime int64 `json:"endTime,omitempty"`
	Limit int `json:"limit,omitempty"`
	Conditions []string `json:"conditions,omitempty"`
	Count bool `json:"count,omitempty"`
	ExcludeOtherAccount bool `json:"excludeOtherAccount,omitempty"`
	Start int `json:"start,omitempty"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
}

// GetAlarmDataParam GetAlarmData request param
type GetAlarmDataParam struct {
	BaseParam
	Params GetAlarmDataDetailParam `json:"params"`
}
