// Copyright (c) ZStack.io, Inc.

package param

// CheckApiPermissionDetailParam CheckApiPermission detail param
type CheckApiPermissionDetailParam struct {
	UserUuid string `json:"userUuid,omitempty"`
	ApiNames []string `json:"apiNames" validate:"required"`
}

// CheckApiPermissionParam CheckApiPermission request param
type CheckApiPermissionParam struct {
	BaseParam
	Params CheckApiPermissionDetailParam `json:"params"`
}
