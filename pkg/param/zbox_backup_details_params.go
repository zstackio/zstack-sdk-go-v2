// Copyright (c) ZStack.io, Inc.

package param

// GetZBoxBackupDetailsDetailParam GetZBoxBackupDetails详细参数
type GetZBoxBackupDetailsDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// GetZBoxBackupDetailsParam GetZBoxBackupDetails请求参数
type GetZBoxBackupDetailsParam struct {
	BaseParam
	Params GetZBoxBackupDetailsDetailParam `json:"params"` // 详细参数
}

