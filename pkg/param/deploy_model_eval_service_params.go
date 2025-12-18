// Copyright (c) ZStack.io, Inc.

package param

// DeployModelEvalServiceDetailParam DeployModelEvalService detail param
type DeployModelEvalServiceDetailParam struct {
	TaskType string `json:"taskType" validate:"required"`
	Limits int `json:"limits" validate:"required"`
	Temperature float32 `json:"temperature,omitempty"`
	TopK int `json:"topK,omitempty"`
	TopP float32 `json:"topP,omitempty"`
	MaxNewTokens int `json:"maxNewTokens,omitempty"`
	RepetitionPenalty float32 `json:"repetitionPenalty,omitempty"`
	MaxLength int `json:"maxLength,omitempty"`
	Prompt string `json:"prompt,omitempty"`
	Model string `json:"model,omitempty"`
	Url string `json:"url,omitempty"`
	Parallel int `json:"parallel,omitempty"`
	LogEveryQuery int `json:"logEveryQuery,omitempty"`
	Api string `json:"api,omitempty"`
	RequestHeaders map[string]string `json:"requestHeaders,omitempty"`
	ConnectTimeout int `json:"connectTimeout,omitempty"`
	ReadTimeout int `json:"readTimeout,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Description string `json:"description,omitempty"`
	ModelUuid string `json:"modelUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	VmImageUuid string `json:"vmImageUuid,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	DatasetUuids []string `json:"datasetUuids,omitempty"`
	ModelServiceGroupUuids []string `json:"modelServiceGroupUuids,omitempty"`
	DockerImage string `json:"dockerImage,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	Name string `json:"name" validate:"required"`
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`
	StartupParameters map[string]string `json:"startupParameters,omitempty"`
	Type string `json:"type" validate:"required"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	L3NetworkUuids []string `json:"l3NetworkUuids,omitempty"`
	ServiceBootUptime int `json:"serviceBootUptime,omitempty"`
	ServiceLivez string `json:"serviceLivez,omitempty"`
	ServiceReadyz string `json:"serviceReadyz,omitempty"`
	RootDiskOfferingUuid string `json:"rootDiskOfferingUuid,omitempty"`
	RootDiskSize int64 `json:"rootDiskSize,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// DeployModelEvalServiceParam DeployModelEvalService request param
type DeployModelEvalServiceParam struct {
	BaseParam
	Params DeployModelEvalServiceDetailParam `json:"params"`
}
