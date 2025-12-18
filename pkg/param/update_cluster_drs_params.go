// Copyright (c) ZStack.io, Inc.

package param

// UpdateClusterDRSDetailParam UpdateClusterDRS detail param
type UpdateClusterDRSDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AutomationLevel string `json:"automationLevel,omitempty"`
	Thresholds []interface{} `json:"thresholds,omitempty"`
	ThresholdDuration int `json:"thresholdDuration,omitempty"`
	State string `json:"state,omitempty"`
}

// UpdateClusterDRSParam UpdateClusterDRS request param
type UpdateClusterDRSParam struct {
	BaseParam
	Params UpdateClusterDRSDetailParam `json:"params"`
}
