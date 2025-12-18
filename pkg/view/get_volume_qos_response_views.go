// Copyright (c) ZStack.io, Inc.

package view

// GetVolumeQosView GetVolumeQos
type GetVolumeQosView struct {
	VolumeUuid string `json:"volumeUuid,omitempty"`
	VolumeBandwidth int64 `json:"volumeBandwidth,omitempty"`
	VolumeBandwidthRead int64 `json:"volumeBandwidthRead,omitempty"`
	VolumeBandwidthWrite int64 `json:"volumeBandwidthWrite,omitempty"`
	IopsTotal int64 `json:"iopsTotal,omitempty"`
	IopsRead int64 `json:"iopsRead,omitempty"`
	IopsWrite int64 `json:"iopsWrite,omitempty"`
	VolumeBandwidthUpthreshold int64 `json:"volumeBandwidthUpthreshold,omitempty"`
	VolumeBandwidthReadUpthreshold int64 `json:"volumeBandwidthReadUpthreshold,omitempty"`
	VolumeBandwidthWriteUpthreshold int64 `json:"volumeBandwidthWriteUpthreshold,omitempty"`
	IopsTotalUpthreshold int64 `json:"iopsTotalUpthreshold,omitempty"`
	IopsReadUpthreshold int64 `json:"iopsReadUpthreshold,omitempty"`
	IopsWriteUpthreshold int64 `json:"iopsWriteUpthreshold,omitempty"`
	Success bool `json:"success,omitempty"`
}

