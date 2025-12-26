// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncImageSize operates on SyncImageSize
func (cli *ZSClient) SyncImageSize(uuid string, params param.SyncImageSizeParam) (*view.SyncImageSizeEventView, error) {
	resp := view.SyncImageSizeEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
