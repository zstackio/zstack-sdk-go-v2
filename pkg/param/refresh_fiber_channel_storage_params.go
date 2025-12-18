// Copyright (c) ZStack.io, Inc.

package param

// RefreshFiberChannelStorageDetailParam RefreshFiberChannelStorage详细参数
type RefreshFiberChannelStorageDetailParam struct {
	rest string `json:"zoneUuid" validate:"required"` // 必填
	rest string `json:"clusterUuid,omitempty"`
	rest []string `json:"scsiLunUuids,omitempty"`
}

// RefreshFiberChannelStorageParam RefreshFiberChannelStorage请求参数
type RefreshFiberChannelStorageParam struct {
	BaseParam
	Params RefreshFiberChannelStorageDetailParam `json:"params"` // 详细参数
}

