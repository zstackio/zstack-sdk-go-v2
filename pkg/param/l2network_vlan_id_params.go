// Copyright (c) ZStack.io, Inc.

package param

// ChangeL2NetworkVlanIdDetailParam ChangeL2NetworkVlanId详细参数
type ChangeL2NetworkVlanIdDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest int `json:"vlan,omitempty"`
	rest string `json:"type,omitempty"`
}

// ChangeL2NetworkVlanIdParam ChangeL2NetworkVlanId请求参数
type ChangeL2NetworkVlanIdParam struct {
	BaseParam
	Params ChangeL2NetworkVlanIdDetailParam `json:"params"` // 详细参数
}

