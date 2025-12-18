// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddThirdpartyPlatform 操作AddThirdpartyPlatform
func (cli *ZSClient) AddThirdpartyPlatform(params param.AddThirdpartyPlatformParam) (*view.AddThirdpartyPlatformEventView, error) {
	resp := view.AddThirdpartyPlatformEventView{}
	if err := cli.Post("v1/zwatch/third-party/platforms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

