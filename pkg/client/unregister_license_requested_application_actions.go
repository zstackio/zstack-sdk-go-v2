// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UnregisterLicenseRequestedApplication 操作UnregisterLicenseRequestedApplication
func (cli *ZSClient) UnregisterLicenseRequestedApplication(uuid string, params param.UnregisterLicenseRequestedApplicationParam) (*view.UnregisterLicenseRequestedApplicationEventView, error) {
	resp := view.UnregisterLicenseRequestedApplicationEventView{}
	if err := cli.Put("v1/license/unregister-applications", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

