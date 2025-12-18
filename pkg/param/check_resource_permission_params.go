// Copyright (c) ZStack.io, Inc.

package param

// CheckResourcePermissionDetailParam CheckResourcePermission detail param
type CheckResourcePermissionDetailParam struct {
	ResourceType string `json:"resourceType" validate:"required"`
}

// CheckResourcePermissionParam CheckResourcePermission request param
type CheckResourcePermissionParam struct {
	BaseParam
	Params CheckResourcePermissionDetailParam `json:"params"`
}
