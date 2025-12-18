// Copyright (c) ZStack.io, Inc.

package param

// DeployDistributedModelServiceDetailParam DeployDistributedModelService详细参数
type DeployDistributedModelServiceDetailParam struct {
	rest []interface{} `json:"modelServices" validate:"required"` // 必填
	rest string `json:"serviceCreationStrategy" validate:"required"` // 必填
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// DeployDistributedModelServiceParam DeployDistributedModelService请求参数
type DeployDistributedModelServiceParam struct {
	BaseParam
	Params DeployDistributedModelServiceDetailParam `json:"params"` // 详细参数
}

