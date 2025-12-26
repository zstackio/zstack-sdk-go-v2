// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateAccessKey creates AccessKey
func (cli *ZSClient) CreateAccessKey(params param.CreateAccessKeyParam) (*view.CreateAccessKeyEventView, error) {
	resp := view.CreateAccessKeyEventView{}
	if err := cli.Post("v1/accesskeys", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
