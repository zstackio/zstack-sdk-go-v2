// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ShareResource 操作ShareResource
func (cli *ZSClient) ShareResource(uuid string, params param.ShareResourceParam) (*view.ShareResourceEventView, error) {
	resp := view.ShareResourceEventView{}
	if err := cli.Put("v1/accounts/resources/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

