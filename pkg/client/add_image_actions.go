// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddImage adds Image
func (cli *ZSClient) AddImage(params param.AddImageParam) (*view.AddImageEventView, error) {
	resp := view.AddImageEventView{}
	if err := cli.Post("v1/images", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
