// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RegisterLicenseServer 操作RegisterLicenseServer
func (cli *ZSClient) RegisterLicenseServer(params param.RegisterLicenseServerParam) (*view.RegisterLicenseServerEventView, error) {
	resp := view.RegisterLicenseServerEventView{}
	if err := cli.Post("v1/license-server/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

