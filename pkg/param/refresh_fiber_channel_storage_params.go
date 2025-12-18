// Copyright (c) ZStack.io, Inc.

package param

// RefreshFiberChannelStorageDetailParam RefreshFiberChannelStorage detail param
type RefreshFiberChannelStorageDetailParam struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	ScsiLunUuids []string `json:"scsiLunUuids,omitempty"`
}

// RefreshFiberChannelStorageParam RefreshFiberChannelStorage request param
type RefreshFiberChannelStorageParam struct {
	BaseParam
	Params RefreshFiberChannelStorageDetailParam `json:"params"`
}
