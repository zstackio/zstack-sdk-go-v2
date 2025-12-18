// Copyright (c) ZStack.io, Inc.

package param

// CreateObservabilityServerDetailParam CreateObservabilityServer详细参数
type CreateObservabilityServerDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"observabilityServerOfferingUuid" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"primaryStorageUuidForRootVolume,omitempty"`
	rest []string `json:"rootVolumeSystemTags,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateObservabilityServerParam CreateObservabilityServer请求参数
type CreateObservabilityServerParam struct {
	BaseParam
	Params CreateObservabilityServerDetailParam `json:"params"` // 详细参数
}

