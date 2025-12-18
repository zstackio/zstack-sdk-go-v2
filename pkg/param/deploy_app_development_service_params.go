// Copyright (c) ZStack.io, Inc.

package param

// DeployAppDevelopmentServiceDetailParam DeployAppDevelopmentService详细参数
type DeployAppDevelopmentServiceDetailParam struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"modelUuid,omitempty"`
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"vmImageUuid,omitempty"`
	rest string `json:"primaryStorageUuid,omitempty"`
	rest []string `json:"datasetUuids,omitempty"`
	rest []string `json:"modelServiceGroupUuids,omitempty"`
	rest string `json:"dockerImage,omitempty"`
	rest int `json:"cpuNum,omitempty"`
	rest string `json:"name" validate:"required"` // 必填
	rest map[string]string `json:"environmentVariables,omitempty"`
	rest map[string]string `json:"startupParameters,omitempty"`
	rest string `json:"type" validate:"required"` // 必填
	rest string `json:"clusterUuid,omitempty"`
	rest int64 `json:"memorySize,omitempty"`
	rest []string `json:"l3NetworkUuids,omitempty"`
	rest int `json:"serviceBootUptime,omitempty"`
	rest string `json:"serviceLivez,omitempty"`
	rest string `json:"serviceReadyz,omitempty"`
	rest string `json:"rootDiskOfferingUuid,omitempty"`
	rest int64 `json:"rootDiskSize,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// DeployAppDevelopmentServiceParam DeployAppDevelopmentService请求参数
type DeployAppDevelopmentServiceParam struct {
	BaseParam
	Params DeployAppDevelopmentServiceDetailParam `json:"params"` // 详细参数
}

