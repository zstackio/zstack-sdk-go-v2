// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CheckIAM2OrganizationAvailability operates on CheckIAM2OrganizationAvailability
func (cli *ZSClient) CheckIAM2OrganizationAvailability(params param.CheckIAM2OrganizationAvailabilityParam) (*view.CheckIAM2OrganizationAvailabilityView, error) {
	var resp view.CheckIAM2OrganizationAvailabilityView
	if err := cli.Get("v1/iam2/organizations/availabilities", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
