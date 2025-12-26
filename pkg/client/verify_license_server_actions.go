// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// VerifyLicenseServer operates on VerifyLicenseServer
func (cli *ZSClient) VerifyLicenseServer(params param.VerifyLicenseServerParam) (*view.VerifyLicenseServerEventView, error) {
	resp := view.VerifyLicenseServerEventView{}
	if err := cli.Post("v1/license-server/register-verify", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
