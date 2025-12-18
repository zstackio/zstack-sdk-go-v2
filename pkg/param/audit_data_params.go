// Copyright (c) ZStack.io, Inc.

package param

// GetAuditDataDetailParam GetAuditData详细参数
type GetAuditDataDetailParam struct {
	rest int64 `json:"startTime,omitempty"`
	rest int64 `json:"endTime,omitempty"`
	rest int `json:"limit,omitempty"`
	rest []string `json:"conditions,omitempty"`
	rest string `json:"auditType,omitempty"`
}

// GetAuditDataParam GetAuditData请求参数
type GetAuditDataParam struct {
	BaseParam
	Params GetAuditDataDetailParam `json:"params"` // 详细参数
}

