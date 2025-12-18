// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateThirdpartyPlatform 更新ThirdpartyPlatform
func (cli *ZSClient) UpdateThirdpartyPlatform(uuid string, params param.UpdateThirdpartyPlatformParam) (*view.UpdateThirdpartyPlatformEventView, error) {
	resp := view.UpdateThirdpartyPlatformEventView{}
	if err := cli.Put("v1/zwatch/third-party/platforms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

