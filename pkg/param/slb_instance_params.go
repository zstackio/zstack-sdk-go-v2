// Copyright (c) ZStack.io, Inc.

package param

// CreateSlbInstanceDetailParam CreateSlbInstance详细参数
type CreateSlbInstanceDetailParam struct {
	rest string `json:"name" validate:"required"` // 必填
	rest string `json:"slbGroupUuid" validate:"required"` // 必填
	rest string `json:"description,omitempty"`
	rest string `json:"zoneUuid,omitempty"`
	rest string `json:"clusterUuid,omitempty"`
	rest string `json:"hostUuid,omitempty"`
	rest string `json:"primaryStorageUuidForRootVolume,omitempty"`
	rest []string `json:"rootVolumeSystemTags,omitempty"`
	rest string `json:"resourceUuid,omitempty"`
	rest []string `json:"tagUuids,omitempty"`
}

// CreateSlbInstanceParam CreateSlbInstance请求参数
type CreateSlbInstanceParam struct {
	BaseParam
	Params CreateSlbInstanceDetailParam `json:"params"` // 详细参数
}

