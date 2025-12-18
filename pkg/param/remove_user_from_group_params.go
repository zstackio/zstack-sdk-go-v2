// Copyright (c) ZStack.io, Inc.

package param

// RemoveUserFromGroupDetailParam RemoveUserFromGroup detail param
type RemoveUserFromGroupDetailParam struct {
	UserUuid string `json:"userUuid" validate:"required"`
	GroupUuid string `json:"groupUuid" validate:"required"`
}

// RemoveUserFromGroupParam RemoveUserFromGroup request param
type RemoveUserFromGroupParam struct {
	BaseParam
	Params RemoveUserFromGroupDetailParam `json:"params"`
}
