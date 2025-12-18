// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateUserTag 创建UserTag
func (cli *ZSClient) CreateUserTag(params param.CreateUserTagParam) (*view.CreateUserTagEventView, error) {
	resp := view.CreateUserTagEventView{}
	if err := cli.Post("v1/user-tags", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

