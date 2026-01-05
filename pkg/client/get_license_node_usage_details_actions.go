// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetLicenseNodeUsageDetails gets LicenseNodeUsageDetails by uuid
func (cli *ZSClient) GetLicenseNodeUsageDetails(uuid string) (*view.GetLicenseNodeUsageDetailsView, error) {
	var resp view.GetLicenseNodeUsageDetailsView
	if err := cli.Get("v1/license/node/usage/details", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
