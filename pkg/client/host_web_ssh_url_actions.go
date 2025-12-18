// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetHostWebSshUrl 获取HostWebSshUrl详情
func (cli *ZSClient) GetHostWebSshUrl(uuid string) (*view.GetHostWebSshUrlEventView, error) {
	var resp view.GetHostWebSshUrlEventView
	if err := cli.Get("v1/hosts/webssh", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

