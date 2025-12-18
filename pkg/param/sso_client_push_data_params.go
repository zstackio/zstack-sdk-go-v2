// Copyright (c) ZStack.io, Inc.

package param

// SsoClientPushDataDetailParam SsoClientPushData详细参数
type SsoClientPushDataDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"dataType,omitempty"`
	rest string `json:"serverUrl,omitempty"`
}

// SsoClientPushDataParam SsoClientPushData请求参数
type SsoClientPushDataParam struct {
	BaseParam
	Params SsoClientPushDataDetailParam `json:"params"` // 详细参数
}

