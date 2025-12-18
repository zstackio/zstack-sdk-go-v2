// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCurrentTime 获取CurrentTime详情
func (cli *ZSClient) GetCurrentTime(uuid string) (*view.GetCurrentTimeView, error) {
	var resp view.GetCurrentTimeView
	if err := cli.Get("v1/management-nodes/actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

