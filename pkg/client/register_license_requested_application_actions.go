// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RegisterLicenseRequestedApplication operates on RegisterLicenseRequestedApplication
func (cli *ZSClient) RegisterLicenseRequestedApplication(params param.RegisterLicenseRequestedApplicationParam) (*view.RegisterLicenseRequestedApplicationEventView, error) {
	resp := view.RegisterLicenseRequestedApplicationEventView{}
	if err := cli.Post("v1/licenses/applications", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
