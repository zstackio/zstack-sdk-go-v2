// Copyright (c) ZStack.io, Inc.

package param

// AttachPoliciesToUserDetailParam AttachPoliciesToUser detail param
type AttachPoliciesToUserDetailParam struct {
	UserUuid string `json:"userUuid" validate:"required"`
	PolicyUuids []string `json:"policyUuids" validate:"required"`
}

// AttachPoliciesToUserParam AttachPoliciesToUser request param
type AttachPoliciesToUserParam struct {
	BaseParam
	Params AttachPoliciesToUserDetailParam `json:"params"`
}
