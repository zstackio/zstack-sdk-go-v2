// Copyright (c) ZStack.io, Inc.

package param

// GetAuditDataDetailParam GetAuditData detail param
type GetAuditDataDetailParam struct {
	StartTime int64 `json:"startTime,omitempty"`
	EndTime int64 `json:"endTime,omitempty"`
	Limit int `json:"limit,omitempty"`
	Conditions []string `json:"conditions,omitempty"`
	AuditType string `json:"auditType,omitempty"`
}

// GetAuditDataParam GetAuditData request param
type GetAuditDataParam struct {
	BaseParam
	Params GetAuditDataDetailParam `json:"params"`
}
