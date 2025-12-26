// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// LoginIAM2Platform operates on LoginIAM2Platform
func (cli *ZSClient) LoginIAM2Platform(uuid string, params param.LoginIAM2PlatformParam) (*view.LoginIAM2PlatformView, error) {
	resp := view.LoginIAM2PlatformView{}
	if err := cli.Put("v1/iam2/platform/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
