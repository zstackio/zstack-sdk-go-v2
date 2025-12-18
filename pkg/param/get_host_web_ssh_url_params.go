// Copyright (c) ZStack.io, Inc.

package param

// GetHostWebSshUrlDetailParam GetHostWebSshUrl detail param
type GetHostWebSshUrlDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Https bool `json:"https,omitempty"`
	UserName string `json:"userName" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// GetHostWebSshUrlParam GetHostWebSshUrl request param
type GetHostWebSshUrlParam struct {
	BaseParam
	Params GetHostWebSshUrlDetailParam `json:"params"`
}
