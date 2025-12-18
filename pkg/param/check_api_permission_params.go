// Copyright (c) ZStack.io, Inc.

package param

// CheckApiPermissionDetailParam CheckApiPermission详细参数
type CheckApiPermissionDetailParam struct {
	rest string `json:"userUuid,omitempty"`
	rest []string `json:"apiNames" validate:"required"` // 必填
}

// CheckApiPermissionParam CheckApiPermission请求参数
type CheckApiPermissionParam struct {
	BaseParam
	Params CheckApiPermissionDetailParam `json:"params"` // 详细参数
}

