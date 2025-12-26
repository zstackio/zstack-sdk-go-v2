// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UnregisterLicenseRequestedApplication operates on UnregisterLicenseRequestedApplication
func (cli *ZSClient) UnregisterLicenseRequestedApplication(uuid string, params param.UnregisterLicenseRequestedApplicationParam) (*view.UnregisterLicenseRequestedApplicationEventView, error) {
	resp := view.UnregisterLicenseRequestedApplicationEventView{}
	if err := cli.Put("v1/license/unregister-applications", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
