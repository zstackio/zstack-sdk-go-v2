// Copyright (c) ZStack.io, Inc.

package param

// QueryDatabaseBackupDetailParam QueryDatabaseBackup详细参数
type QueryDatabaseBackupDetailParam struct {
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

// QueryDatabaseBackupParam QueryDatabaseBackup请求参数
type QueryDatabaseBackupParam struct {
	BaseParam
	Params QueryDatabaseBackupDetailParam `json:"params"` // 详细参数
}

