// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetLicenseInfo gets LicenseInfo by uuid
func (cli *ZSClient) GetLicenseInfo(uuid string) (*view.GetLicenseInfoView, error) {
	var resp view.GetLicenseInfoView
	if err := cli.Get("v1/licenses", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
