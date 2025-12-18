// Copyright (c) ZStack.io, Inc.

package param

// GetInvocationRecordsDetailParam GetInvocationRecords detail param
type GetInvocationRecordsDetailParam struct {
	RecordUuid string `json:"recordUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	IncludeOutput bool `json:"includeOutput,omitempty"`
	Limit int `json:"limit,omitempty"`
	Start int `json:"start,omitempty"`
}

// GetInvocationRecordsParam GetInvocationRecords request param
type GetInvocationRecordsParam struct {
	BaseParam
	Params GetInvocationRecordsDetailParam `json:"params"`
}
