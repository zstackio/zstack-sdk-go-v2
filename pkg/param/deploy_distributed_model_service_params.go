// Copyright (c) ZStack.io, Inc.

package param

// DeployDistributedModelServiceDetailParam DeployDistributedModelService detail param
type DeployDistributedModelServiceDetailParam struct {
	ModelServices []ModelServiceParam `json:"modelServices" validate:"required"`
	ServiceCreationStrategy string `json:"serviceCreationStrategy" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// DeployDistributedModelServiceParam DeployDistributedModelService request param
type DeployDistributedModelServiceParam struct {
	BaseParam
	Params DeployDistributedModelServiceDetailParam `json:"params"`
}
