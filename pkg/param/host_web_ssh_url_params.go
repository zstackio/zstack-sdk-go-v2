// Copyright (c) ZStack.io, Inc.

package param

// GetHostWebSshUrlDetailParam GetHostWebSshUrl详细参数
type GetHostWebSshUrlDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest bool `json:"https,omitempty"`
	rest string `json:"userName" validate:"required"` // 必填
	rest string `json:"password" validate:"required"` // 必填
}

// GetHostWebSshUrlParam GetHostWebSshUrl请求参数
type GetHostWebSshUrlParam struct {
	BaseParam
	Params GetHostWebSshUrlDetailParam `json:"params"` // 详细参数
}

