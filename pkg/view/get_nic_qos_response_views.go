// Copyright (c) ZStack.io, Inc.

package view

// GetNicQosView GetNicQos
type GetNicQosView struct {
	OutboundBandwidth int64 `json:"outboundBandwidth,omitempty"`
	InboundBandwidth int64 `json:"inboundBandwidth,omitempty"`
	OutboundBandwidthUpthreshold int64 `json:"outboundBandwidthUpthreshold,omitempty"`
	InboundBandwidthUpthreshold int64 `json:"inboundBandwidthUpthreshold,omitempty"`
	Success bool `json:"success,omitempty"`
}

