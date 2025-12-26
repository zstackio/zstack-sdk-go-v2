// Copyright (c) ZStack.io, Inc.

package view

// GetZWatchAlertHistogramView GetZWatchAlertHistogram
type GetZWatchAlertHistogramView struct {
	Histograms []HistogramView `json:"histograms,omitempty"`
	Success bool `json:"success,omitempty"`
}

