// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetLicenseAuthorizedCapacity gets LicenseAuthorizedCapacity by uuid
func (cli *ZSClient) GetLicenseAuthorizedCapacity(uuid string) (*view.GetLicenseAuthorizedCapacityView, error) {
	var resp view.GetLicenseAuthorizedCapacityView
	if err := cli.Get("v1/license-server/authorized-capacity", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
