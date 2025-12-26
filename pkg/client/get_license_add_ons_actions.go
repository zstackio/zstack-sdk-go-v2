// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetLicenseAddOns gets LicenseAddOns by uuid
func (cli *ZSClient) GetLicenseAddOns(uuid string) (*view.GetLicenseAddOnsView, error) {
	var resp view.GetLicenseAddOnsView
	if err := cli.Get("v1/licenses/addons", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
