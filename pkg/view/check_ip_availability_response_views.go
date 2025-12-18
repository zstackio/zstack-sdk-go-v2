// Copyright (c) ZStack.io, Inc.

package view

// CheckIpAvailabilityView CheckIpAvailability
type CheckIpAvailabilityView struct {
	Available bool `json:"available,omitempty"`
	Reason string `json:"reason,omitempty"`
	Success bool `json:"success,omitempty"`
}

