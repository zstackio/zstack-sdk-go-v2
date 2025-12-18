// Copyright (c) ZStack.io, Inc.

package param

// QueryGuestVmScriptExecutedRecordDetailParam QueryGuestVmScriptExecutedRecord详细参数
type QueryGuestVmScriptExecutedRecordDetailParam struct {
	rest []interface{} `json:"conditions" validate:"required"` // 必填
	rest int `json:"limit,omitempty"`
	rest int `json:"start,omitempty"`
	rest bool `json:"count,omitempty"`
	rest string `json:"groupBy,omitempty"`
	rest bool `json:"replyWithCount,omitempty"`
	rest string `json:"filterName,omitempty"`
	rest string `json:"sortBy,omitempty"`
	rest string `json:"sortDirection,omitempty"`
	rest []string `json:"fields,omitempty"`
}

// QueryGuestVmScriptExecutedRecordParam QueryGuestVmScriptExecutedRecord请求参数
type QueryGuestVmScriptExecutedRecordParam struct {
	BaseParam
	Params QueryGuestVmScriptExecutedRecordDetailParam `json:"params"` // 详细参数
}

