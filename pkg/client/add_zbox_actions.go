// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddZBox adds ZBox
func (cli *ZSClient) AddZBox(params param.AddZBoxParam) (*view.AddZBoxEventView, error) {
	resp := view.AddZBoxEventView{}
	if err := cli.Post("v1/zbox", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
