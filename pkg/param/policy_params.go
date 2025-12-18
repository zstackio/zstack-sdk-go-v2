// Copyright (c) ZStack.io, Inc.

package param

// DeletePolicyDetailParam DeletePolicy详细参数
type DeletePolicyDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"deleteMode,omitempty"`
}

// DeletePolicyParam DeletePolicy请求参数
type DeletePolicyParam struct {
	BaseParam
	Params DeletePolicyDetailParam `json:"params"` // 详细参数
}

