// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeResourceOwner 操作ResourceOwner
func (cli *ZSClient) ChangeResourceOwner(params param.ChangeResourceOwnerParam) (*view.ChangeResourceOwnerEventView, error) {
	resp := view.ChangeResourceOwnerEventView{}
	if err := cli.Post("v1/account/{accountUuid}/resources", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

