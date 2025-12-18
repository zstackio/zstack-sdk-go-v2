// Copyright (c) ZStack.io, Inc.

package param

// RecoverResourceSplitBrainDetailParam RecoverResourceSplitBrain detail param
type RecoverResourceSplitBrainDetailParam struct {
	ResourceUuid string `json:"resourceUuid" validate:"required"`
	HostUuid string `json:"hostUuid" validate:"required"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	ForceRecover bool `json:"forceRecover,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
}

// RecoverResourceSplitBrainParam RecoverResourceSplitBrain request param
type RecoverResourceSplitBrainParam struct {
	BaseParam
	Params RecoverResourceSplitBrainDetailParam `json:"params"`
}
