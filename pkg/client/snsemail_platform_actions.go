// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSNSEmailPlatform 创建SNSEmailPlatform
func (cli *ZSClient) CreateSNSEmailPlatform(params param.CreateSNSEmailPlatformParam) (*view.CreateSNSApplicationPlatformEventView, error) {
	resp := view.CreateSNSApplicationPlatformEventView{}
	if err := cli.Post("v1/sns/application-platforms/email", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

