// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now // avoid unused import

// DeleteClusterDRSParamDetail DeleteClusterDRS detail param
type DeleteClusterDRSParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	DeleteMode string `json:"deleteMode,omitempty"`
}

// DeleteClusterDRSParam DeleteClusterDRS request param
type DeleteClusterDRSParam struct {
	BaseParam
	DeleteClusterDRS DeleteClusterDRSParamDetail `json:"deleteClusterDRS"`
}
// CreateClusterDRSParamDetail CreateClusterDRS detail param
type CreateClusterDRSParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
	AutomationLevel string `json:"automationLevel" validate:"required"`
	Thresholds []ThresholdParam `json:"thresholds" validate:"required"`
	ThresholdDuration int `json:"thresholdDuration" validate:"required"`
	DefaultEnable bool `json:"defaultEnable,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateClusterDRSParam CreateClusterDRS request param
type CreateClusterDRSParam struct {
	BaseParam
	CreateClusterDRS CreateClusterDRSParamDetail `json:"createClusterDRS"`
}
// UpdateClusterDRSParamDetail UpdateClusterDRS detail param
type UpdateClusterDRSParamDetail struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	AutomationLevel string `json:"automationLevel,omitempty"`
	Thresholds []ThresholdParam `json:"thresholds,omitempty"`
	ThresholdDuration int `json:"thresholdDuration,omitempty"`
	State string `json:"state,omitempty"`
}

// UpdateClusterDRSParam UpdateClusterDRS request param
type UpdateClusterDRSParam struct {
	BaseParam
	UpdateClusterDRS UpdateClusterDRSParamDetail `json:"updateClusterDRS"`
}
