// Copyright (c) ZStack.io, Inc.

package param

// UpdateThirdpartyAlertsDetailParam UpdateThirdpartyAlerts detail param
type UpdateThirdpartyAlertsDetailParam struct {
	Uuid string `json:"uuid,omitempty"`
	StartTimeMillis int64 `json:"startTimeMillis,omitempty"`
	EndTimeMillis int64 `json:"endTimeMillis,omitempty"`
	UpdateReadStatus string `json:"updateReadStatus,omitempty"`
}

// UpdateThirdpartyAlertsParam UpdateThirdpartyAlerts request param
type UpdateThirdpartyAlertsParam struct {
	BaseParam
	Params UpdateThirdpartyAlertsDetailParam `json:"params"`
}
