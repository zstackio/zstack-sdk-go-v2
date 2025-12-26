// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetLicenseUKeyStatus gets LicenseUKeyStatus by uuid
func (cli *ZSClient) GetLicenseUKeyStatus(uuid string) (*view.GetLicenseUKeyStatusEventView, error) {
	var resp view.GetLicenseUKeyStatusEventView
	if err := cli.Get("v1/licenses/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
