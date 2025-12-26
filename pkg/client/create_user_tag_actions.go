// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateUserTag creates UserTag
func (cli *ZSClient) CreateUserTag(params param.CreateUserTagParam) (*view.CreateUserTagEventView, error) {
	resp := view.CreateUserTagEventView{}
	if err := cli.Post("v1/user-tags", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
