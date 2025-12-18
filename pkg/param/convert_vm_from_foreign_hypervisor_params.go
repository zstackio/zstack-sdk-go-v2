// Copyright (c) ZStack.io, Inc.

package param

// ConvertVmFromForeignHypervisorDetailParam ConvertVmFromForeignHypervisor detail param
type ConvertVmFromForeignHypervisorDetailParam struct {
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ConversionHostUuid string `json:"conversionHostUuid,omitempty"`
	SshPrivKey string `json:"sshPrivKey,omitempty"`
	CpuNum int `json:"cpuNum" validate:"required"`
	MemorySize int64 `json:"memorySize" validate:"required"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid" validate:"required"`
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty"`
	Platform string `json:"platform,omitempty"`
	Type string `json:"type,omitempty"`
	Strategy string `json:"strategy,omitempty"`
	ConvertStrategy string `json:"convertStrategy,omitempty"`
	PauseVm bool `json:"pauseVm,omitempty"`
	VolumeFilters []interface{} `json:"volumeFilters,omitempty"`
	RootFileSystem string `json:"rootFileSystem,omitempty"`
	LongJobName string `json:"longJobName,omitempty"`
	LongJobDescription string `json:"longJobDescription,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// ConvertVmFromForeignHypervisorParam ConvertVmFromForeignHypervisor request param
type ConvertVmFromForeignHypervisorParam struct {
	BaseParam
	Params ConvertVmFromForeignHypervisorDetailParam `json:"params"`
}
