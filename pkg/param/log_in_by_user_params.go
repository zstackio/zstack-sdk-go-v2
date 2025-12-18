// Copyright (c) ZStack.io, Inc.

package param

// LogInByUserDetailParam LogInByUser detail param
type LogInByUserDetailParam struct {
	AccountUuid string `json:"accountUuid,omitempty"`
	AccountName string `json:"accountName,omitempty"`
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
	ClientInfo map[string]string `json:"clientInfo,omitempty"`
}

// LogInByUserParam LogInByUser request param
type LogInByUserParam struct {
	BaseParam
	Params LogInByUserDetailParam `json:"params"`
}
