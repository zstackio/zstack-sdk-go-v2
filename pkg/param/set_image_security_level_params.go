// Copyright (c) ZStack.io, Inc.

package param

// SetImageSecurityLevelDetailParam SetImageSecurityLevel详细参数
type SetImageSecurityLevelDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"securityLevel,omitempty"`
}

// SetImageSecurityLevelParam SetImageSecurityLevel请求参数
type SetImageSecurityLevelParam struct {
	BaseParam
	Params SetImageSecurityLevelDetailParam `json:"params"` // 详细参数
}

