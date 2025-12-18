// Copyright (c) ZStack.io, Inc.

package param

// GetInvocationRecordsDetailParam GetInvocationRecords详细参数
type GetInvocationRecordsDetailParam struct {
	rest string `json:"recordUuid,omitempty"`
	rest string `json:"vmInstanceUuid,omitempty"`
	rest bool `json:"includeOutput,omitempty"`
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
}

// GetInvocationRecordsParam GetInvocationRecords请求参数
type GetInvocationRecordsParam struct {
	BaseParam
	Params GetInvocationRecordsDetailParam `json:"params"` // 详细参数
}

