// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetLicenseAuthorizedCapacity gets LicenseAuthorizedCapacity by uuid
func (cli *ZSClient) GetLicenseAuthorizedCapacity(uuid string) (*view.GetLicenseAuthorizedCapacityView, error) {
	var resp view.GetLicenseAuthorizedCapacityView
	if err := cli.Get("v1/license-server/authorized-capacity", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
