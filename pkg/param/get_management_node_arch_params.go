// Copyright (c) ZStack.io, Inc.

package param

// GetManagementNodeArchDetailParam GetManagementNodeArch detail param
type GetManagementNodeArchDetailParam struct {
}

// GetManagementNodeArchParam GetManagementNodeArch request param
type GetManagementNodeArchParam struct {
	BaseParam
	Params GetManagementNodeArchDetailParam `json:"params"`
}
