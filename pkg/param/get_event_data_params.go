// Copyright (c) ZStack.io, Inc.

package param

// GetEventDataDetailParam GetEventData detail param
type GetEventDataDetailParam struct {
	StartTime int64 `json:"startTime,omitempty"`
	EndTime int64 `json:"endTime,omitempty"`
	OffsetAheadOfCurrentTime int64 `json:"offsetAheadOfCurrentTime,omitempty"`
	Limit int `json:"limit,omitempty"`
	Conditions []string `json:"conditions,omitempty"`
	Count bool `json:"count,omitempty"`
	Start int `json:"start,omitempty"`
	ConditionExpression string `json:"conditionExpression,omitempty"`
	EndpointUuid string `json:"endpointUuid,omitempty"`
}

// GetEventDataParam GetEventData request param
type GetEventDataParam struct {
	BaseParam
	Params GetEventDataDetailParam `json:"params"`
}
