// Copyright (c) ZStack.io, Inc.

package param

// AttachProvisionNicToBondingDetailParam AttachProvisionNicToBonding详细参数
type AttachProvisionNicToBondingDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"provisionNicUuid" validate:"required"` // 必填
	rest string `json:"bondingUuid" validate:"required"` // 必填
	rest string `json:"provisionIp,omitempty"`
	rest string `json:"customMac,omitempty"`
}

// AttachProvisionNicToBondingParam AttachProvisionNicToBonding请求参数
type AttachProvisionNicToBondingParam struct {
	BaseParam
	Params AttachProvisionNicToBondingDetailParam `json:"params"` // 详细参数
}

