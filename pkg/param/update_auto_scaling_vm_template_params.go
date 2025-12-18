// Copyright (c) ZStack.io, Inc.

package param

// UpdateAutoScalingVmTemplateDetailParam UpdateAutoScalingVmTemplate detail param
type UpdateAutoScalingVmTemplateDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	VmInstanceName string `json:"vmInstanceName,omitempty"`
	VmInstanceDescription string `json:"vmInstanceDescription,omitempty"`
	VmInstanceOfferingUuid string `json:"vmInstanceOfferingUuid,omitempty"`
	ImageUuid string `json:"imageUuid,omitempty"`
	VmInstanceClusterUuid string `json:"vmInstanceClusterUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
}

// UpdateAutoScalingVmTemplateParam UpdateAutoScalingVmTemplate request param
type UpdateAutoScalingVmTemplateParam struct {
	BaseParam
	Params UpdateAutoScalingVmTemplateDetailParam `json:"params"`
}
