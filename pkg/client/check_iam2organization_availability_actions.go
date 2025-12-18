// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckIAM2OrganizationAvailability 操作CheckIAM2OrganizationAvailability
func (cli *ZSClient) CheckIAM2OrganizationAvailability(params param.CheckIAM2OrganizationAvailabilityParam) (*view.CheckIAM2OrganizationAvailabilityView, error) {
	var resp view.CheckIAM2OrganizationAvailabilityView
	if err := cli.Get("v1/iam2/organizations/availabilities", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

