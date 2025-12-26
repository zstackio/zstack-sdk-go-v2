// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateThirdpartyPlatform updates ThirdpartyPlatform
func (cli *ZSClient) UpdateThirdpartyPlatform(uuid string, params param.UpdateThirdpartyPlatformParam) (*view.UpdateThirdpartyPlatformEventView, error) {
	resp := view.UpdateThirdpartyPlatformEventView{}
	if err := cli.Put("v1/zwatch/third-party/platforms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
