// Copyright (c) ZStack.io, Inc.

package param

// CreateSNSSnmpPlatformDetailParam CreateSNSSnmpPlatform详细参数
type CreateSNSSnmpPlatformDetailParam struct {
	rest string `json:"snmpAddress" validate:"required"` // 必填
	rest int `json:"snmpPort" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateSNSSnmpPlatformParam CreateSNSSnmpPlatform请求参数
type CreateSNSSnmpPlatformParam struct {
	BaseParam
	Params CreateSNSSnmpPlatformDetailParam `json:"params"` // 详细参数
}

