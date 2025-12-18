// Copyright (c) ZStack.io, Inc.

package param

// CreateCdpPolicyDetailParam CreateCdpPolicy detail param
type CreateCdpPolicyDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	RetentionTimePerDay int `json:"retentionTimePerDay,omitempty"`
	HourlyRpSinceDay int `json:"hourlyRpSinceDay,omitempty"`
	DailyRpSinceDay int `json:"dailyRpSinceDay,omitempty"`
	ExpireTimeInDay int `json:"expireTimeInDay,omitempty"`
	FullBackupIntervalInDay int `json:"fullBackupIntervalInDay,omitempty"`
	RecoveryPointPerSecond int `json:"recoveryPointPerSecond" validate:"required"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateCdpPolicyParam CreateCdpPolicy request param
type CreateCdpPolicyParam struct {
	BaseParam
	Params CreateCdpPolicyDetailParam `json:"params"`
}
