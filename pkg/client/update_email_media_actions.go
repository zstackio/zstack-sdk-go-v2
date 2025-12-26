// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateEmailMedia updates EmailMedia
func (cli *ZSClient) UpdateEmailMedia(uuid string, params param.UpdateEmailMediaParam) (*view.UpdateEmailMediaEventView, error) {
	resp := view.UpdateEmailMediaEventView{}
	if err := cli.Put("v1/media/emails/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
