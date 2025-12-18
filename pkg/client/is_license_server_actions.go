// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// IsLicenseServer operates on IsLicenseServer
func (cli *ZSClient) IsLicenseServer(params param.IsLicenseServerParam) (*view.IsLicenseServerView, error) {
	var resp view.IsLicenseServerView
	if err := cli.Get("v1/license-server/is-server", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
