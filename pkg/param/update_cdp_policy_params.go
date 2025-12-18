// Copyright (c) ZStack.io, Inc.

package param

// UpdateCdpPolicyDetailParam UpdateCdpPolicy detail param
type UpdateCdpPolicyDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	RetentionTimePerDay int `json:"retentionTimePerDay,omitempty"`
	HourlyRpSinceDay int `json:"hourlyRpSinceDay,omitempty"`
	DailyRpSinceDay int `json:"dailyRpSinceDay,omitempty"`
	ExpireTimeInDay int `json:"expireTimeInDay,omitempty"`
	FullBackupIntervalInDay int `json:"fullBackupIntervalInDay,omitempty"`
	RecoveryPointPerSecond int `json:"recoveryPointPerSecond,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// UpdateCdpPolicyParam UpdateCdpPolicy request param
type UpdateCdpPolicyParam struct {
	BaseParam
	Params UpdateCdpPolicyDetailParam `json:"params"`
}
