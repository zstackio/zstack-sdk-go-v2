// Copyright (c) ZStack.io, Inc.

package param

// GetManagementNodeOSDetailParam GetManagementNodeOS detail param
type GetManagementNodeOSDetailParam struct {
}

// GetManagementNodeOSParam GetManagementNodeOS request param
type GetManagementNodeOSParam struct {
	BaseParam
	Params GetManagementNodeOSDetailParam `json:"params"`
}
