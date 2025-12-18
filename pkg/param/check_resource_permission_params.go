// Copyright (c) ZStack.io, Inc.

package param

// CheckResourcePermissionDetailParam CheckResourcePermission详细参数
type CheckResourcePermissionDetailParam struct {
	rest string `json:"resourceType" validate:"required"` // 必填
}

// CheckResourcePermissionParam CheckResourcePermission请求参数
type CheckResourcePermissionParam struct {
	BaseParam
	Params CheckResourcePermissionDetailParam `json:"params"` // 详细参数
}

