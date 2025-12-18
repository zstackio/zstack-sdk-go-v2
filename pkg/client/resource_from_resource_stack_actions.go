// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetResourceFromResourceStack 获取ResourceFromResourceStack详情
func (cli *ZSClient) GetResourceFromResourceStack(uuid string) (*view.GetResourceFromResourceStackView, error) {
	var resp view.GetResourceFromResourceStackView
	if err := cli.Get("v1/cloudformation/stack/resources", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

