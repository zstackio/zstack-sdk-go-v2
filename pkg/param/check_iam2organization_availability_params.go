// Copyright (c) ZStack.io, Inc.

package param

// CheckIAM2OrganizationAvailabilityDetailParam CheckIAM2OrganizationAvailability详细参数
type CheckIAM2OrganizationAvailabilityDetailParam struct {
}

// CheckIAM2OrganizationAvailabilityParam CheckIAM2OrganizationAvailability请求参数
type CheckIAM2OrganizationAvailabilityParam struct {
	BaseParam
	Params CheckIAM2OrganizationAvailabilityDetailParam `json:"params"` // 详细参数
}

