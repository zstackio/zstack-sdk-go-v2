// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RegisterLicenseServer operates on RegisterLicenseServer
func (cli *ZSClient) RegisterLicenseServer(params param.RegisterLicenseServerParam) (*view.RegisterLicenseServerEventView, error) {
	resp := view.RegisterLicenseServerEventView{}
	if err := cli.Post("v1/license-server/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
