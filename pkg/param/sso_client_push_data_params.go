// Copyright (c) ZStack.io, Inc.

package param

// SsoClientPushDataDetailParam SsoClientPushData detail param
type SsoClientPushDataDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	DataType string `json:"dataType,omitempty"`
	ServerUrl string `json:"serverUrl,omitempty"`
}

// SsoClientPushDataParam SsoClientPushData request param
type SsoClientPushDataParam struct {
	BaseParam
	Params SsoClientPushDataDetailParam `json:"params"`
}
