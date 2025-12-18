// Copyright (c) ZStack.io, Inc.

package param

// UpdateCdpPolicyDetailParam UpdateCdpPolicy详细参数
type UpdateCdpPolicyDetailParam struct {
	rest string `json:"uuid" validate:"required"` // 必填
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest int `json:"retentionTimePerDay,omitempty"`
	rest int `json:"hourlyRpSinceDay,omitempty"`
	rest int `json:"dailyRpSinceDay,omitempty"`
	rest int `json:"expireTimeInDay,omitempty"`
	rest int `json:"fullBackupIntervalInDay,omitempty"`
	rest int `json:"recoveryPointPerSecond,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// UpdateCdpPolicyParam UpdateCdpPolicy请求参数
type UpdateCdpPolicyParam struct {
	BaseParam
	Params UpdateCdpPolicyDetailParam `json:"params"` // 详细参数
}

