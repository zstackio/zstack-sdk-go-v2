// Copyright (c) ZStack.io, Inc.

package param

// ProvisionSlbInstanceDetailParam ProvisionSlbInstance详细参数
type ProvisionSlbInstanceDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
}

// ProvisionSlbInstanceParam ProvisionSlbInstance请求参数
type ProvisionSlbInstanceParam struct {
	BaseParam
	Params ProvisionSlbInstanceDetailParam `json:"params"` // 详细参数
}

