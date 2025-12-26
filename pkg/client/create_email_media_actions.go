// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateEmailMedia creates EmailMedia
func (cli *ZSClient) CreateEmailMedia(params param.CreateEmailMediaParam) (*view.CreateMediaEventView, error) {
	resp := view.CreateMediaEventView{}
	if err := cli.Post("v1/media/emails", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
