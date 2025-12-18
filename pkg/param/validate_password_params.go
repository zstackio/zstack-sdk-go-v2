// Copyright (c) ZStack.io, Inc.

package param

// ValidatePasswordDetailParam ValidatePassword detail param
type ValidatePasswordDetailParam struct {
	LoginName string `json:"loginName" validate:"required"`
	Password string `json:"password" validate:"required"`
	LoginType string `json:"loginType" validate:"required"`
}

// ValidatePasswordParam ValidatePassword request param
type ValidatePasswordParam struct {
	BaseParam
	Params ValidatePasswordDetailParam `json:"params"`
}
