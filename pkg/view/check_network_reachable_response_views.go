// Copyright (c) ZStack.io, Inc.

package view

// CheckNetworkReachableView CheckNetworkReachable
type CheckNetworkReachableView struct {
	Results []NetworkReachablePairView `json:"results,omitempty"`
	Success bool `json:"success,omitempty"`
}

