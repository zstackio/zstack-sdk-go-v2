// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddThirdpartyPlatform adds ThirdpartyPlatform
func (cli *ZSClient) AddThirdpartyPlatform(params param.AddThirdpartyPlatformParam) (*view.AddThirdpartyPlatformEventView, error) {
	resp := view.AddThirdpartyPlatformEventView{}
	if err := cli.Post("v1/zwatch/third-party/platforms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
