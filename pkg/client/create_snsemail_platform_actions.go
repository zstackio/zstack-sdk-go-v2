// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateSNSEmailPlatform creates SNSEmailPlatform
func (cli *ZSClient) CreateSNSEmailPlatform(params param.CreateSNSEmailPlatformParam) (*view.CreateSNSApplicationPlatformEventView, error) {
	resp := view.CreateSNSApplicationPlatformEventView{}
	if err := cli.Post("v1/sns/application-platforms/email", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
