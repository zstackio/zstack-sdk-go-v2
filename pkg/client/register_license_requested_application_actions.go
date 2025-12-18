// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RegisterLicenseRequestedApplication 操作RegisterLicenseRequestedApplication
func (cli *ZSClient) RegisterLicenseRequestedApplication(params param.RegisterLicenseRequestedApplicationParam) (*view.RegisterLicenseRequestedApplicationEventView, error) {
	resp := view.RegisterLicenseRequestedApplicationEventView{}
	if err := cli.Post("v1/licenses/applications", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

