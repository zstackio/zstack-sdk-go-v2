// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// LoginIAM2Platform 操作LoginIAM2Platform
func (cli *ZSClient) LoginIAM2Platform(uuid string, params param.LoginIAM2PlatformParam) (*view.LoginIAM2PlatformView, error) {
	resp := view.LoginIAM2PlatformView{}
	if err := cli.Put("v1/iam2/platform/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

