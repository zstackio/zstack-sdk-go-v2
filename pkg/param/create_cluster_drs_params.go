// Copyright (c) ZStack.io, Inc.

package param

// CreateClusterDRSDetailParam CreateClusterDRS detail param
type CreateClusterDRSDetailParam struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	AutomationLevel string `json:"automationLevel" validate:"required"`
	Thresholds []interface{} `json:"thresholds" validate:"required"`
	ThresholdDuration int `json:"thresholdDuration" validate:"required"`
	DefaultEnable bool `json:"defaultEnable,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateClusterDRSParam CreateClusterDRS request param
type CreateClusterDRSParam struct {
	BaseParam
	Params CreateClusterDRSDetailParam `json:"params"`
}
