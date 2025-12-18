// Copyright (c) ZStack.io, Inc.

package param

// CheckStaticProvisionIpDetailParam CheckStaticProvisionIp详细参数
type CheckStaticProvisionIpDetailParam struct {
	rest string `json:"clusterUuid" validate:"required"` // 必填
	rest string `json:"provisionIp" validate:"required"` // 必填
}

// CheckStaticProvisionIpParam CheckStaticProvisionIp请求参数
type CheckStaticProvisionIpParam struct {
	BaseParam
	Params CheckStaticProvisionIpDetailParam `json:"params"` // 详细参数
}

