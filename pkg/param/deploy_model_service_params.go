// Copyright (c) ZStack.io, Inc.

package param

// DeployModelServiceDetailParam DeployModelService detail param
type DeployModelServiceDetailParam struct {
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

// DeployModelServiceParam DeployModelService request param
type DeployModelServiceParam struct {
	BaseParam
	Params DeployModelServiceDetailParam `json:"params"`
}
