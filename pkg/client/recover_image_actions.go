// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RecoverImage operates on Image
func (cli *ZSClient) RecoverImage(uuid string, params param.RecoverImageParam) (*view.RecoverImageEventView, error) {
	resp := view.RecoverImageEventView{}
	if err := cli.Put("v1/images/{imageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
