// Copyright (c) ZStack.io, Inc.

package param

// DeleteUserProxyConfigDetailParam DeleteUserProxyConfig detail param
type DeleteUserProxyConfigDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
}

// DeleteUserProxyConfigParam DeleteUserProxyConfig request param
type DeleteUserProxyConfigParam struct {
	BaseParam
	Params DeleteUserProxyConfigDetailParam `json:"params"`
}
