// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetLicenseCapabilities gets LicenseCapabilities by uuid
func (cli *ZSClient) GetLicenseCapabilities(uuid string) (*view.GetLicenseCapabilitiesView, error) {
	var resp view.GetLicenseCapabilitiesView
	if err := cli.Get("v1/licenses/capabilities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
