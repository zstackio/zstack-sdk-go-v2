// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpgradeToLicenseServer operates on UpgradeToLicenseServer
func (cli *ZSClient) UpgradeToLicenseServer(params param.UpgradeToLicenseServerParam) (*view.UpgradeToLicenseServerEventView, error) {
	resp := view.UpgradeToLicenseServerEventView{}
	if err := cli.Post("v1/license-server", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
