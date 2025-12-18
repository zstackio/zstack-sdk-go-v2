// Copyright (c) ZStack.io, Inc.

package param

// CreateAutoScalingVmTemplateDetailParam CreateAutoScalingVmTemplate detail param
type CreateAutoScalingVmTemplateDetailParam struct {
	VmInstanceName string `json:"vmInstanceName" validate:"required"`
	VmInstanceDescription string `json:"vmInstanceDescription,omitempty"`
	VmInstanceOfferingUuid string `json:"vmInstanceOfferingUuid" validate:"required"`
	ImageUuid string `json:"imageUuid" validate:"required"`
	L3NetworkUuids []string `json:"l3NetworkUuids" validate:"required"`
	VmInstanceType string `json:"vmInstanceType,omitempty"`
	RootDiskOfferingUuid string `json:"rootDiskOfferingUuid,omitempty"`
	DataDiskOfferingUuids []string `json:"dataDiskOfferingUuids,omitempty"`
	VmInstanceZoneUuid string `json:"vmInstanceZoneUuid,omitempty"`
	VmInstanceClusterUuid string `json:"vmInstanceClusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	PrimaryStorageUuidForRootVolume string `json:"primaryStorageUuidForRootVolume,omitempty"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid" validate:"required"`
	Strategy string `json:"strategy,omitempty"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateAutoScalingVmTemplateParam CreateAutoScalingVmTemplate request param
type CreateAutoScalingVmTemplateParam struct {
	BaseParam
	Params CreateAutoScalingVmTemplateDetailParam `json:"params"`
}
