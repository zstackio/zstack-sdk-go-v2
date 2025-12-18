// Copyright (c) ZStack.io, Inc.

package param

// CheckIAM2OrganizationAvailabilityDetailParam CheckIAM2OrganizationAvailability detail param
type CheckIAM2OrganizationAvailabilityDetailParam struct {
}

// CheckIAM2OrganizationAvailabilityParam CheckIAM2OrganizationAvailability request param
type CheckIAM2OrganizationAvailabilityParam struct {
	BaseParam
	Params CheckIAM2OrganizationAvailabilityDetailParam `json:"params"`
}
