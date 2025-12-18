// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetSupportAPIs 获取SupportAPIs详情
func (cli *ZSClient) GetSupportAPIs(uuid string) (*view.GetSupportAPIsView, error) {
	var resp view.GetSupportAPIsView
	if err := cli.Get("v1/management-nodes/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

