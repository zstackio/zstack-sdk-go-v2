// Copyright (c) ZStack.io, Inc.

package view

import "time"

// CdpPolicyInventoryView CdpPolicy
type CdpPolicyInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"description,omitempty"`
	rest int `json:"retentionTimePerDay,omitempty"`
	rest int `json:"hourlyRpSinceDay,omitempty"`
	rest int `json:"dailyRpSinceDay,omitempty"`
	rest int `json:"expireTimeInDay,omitempty"`
	rest int `json:"fullBackupIntervalInDay,omitempty"`
	rest int `json:"recoveryPointPerSecond,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

