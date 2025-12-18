// Copyright (c) ZStack.io, Inc.

package param

// DeployModelEvalServiceDetailParam DeployModelEvalService详细参数
type DeployModelEvalServiceDetailParam struct {
	rest string `json:"taskType" validate:"required"` // 必填
	rest int `json:"limits" validate:"required"` // 必填
	rest float32 `json:"temperature,omitempty"`
	rest int `json:"topK,omitempty"`
	rest float32 `json:"topP,omitempty"`
	rest int `json:"maxNewTokens,omitempty"`
	rest float32 `json:"repetitionPenalty,omitempty"`
	rest int `json:"maxLength,omitempty"`
	rest string `json:"prompt,omitempty"`
	rest string `json:"model,omitempty"`
	rest string `json:"url,omitempty"`
	rest int `json:"parallel,omitempty"`
	rest int `json:"logEveryQuery,omitempty"`
	rest string `json:"api,omitempty"`
	rest map[string]string `json:"requestHeaders,omitempty"`
	rest int `json:"connectTimeout,omitempty"`
	rest int `json:"readTimeout,omitempty"`
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

// DeployModelEvalServiceParam DeployModelEvalService请求参数
type DeployModelEvalServiceParam struct {
	BaseParam
	Params DeployModelEvalServiceDetailParam `json:"params"` // 详细参数
}

