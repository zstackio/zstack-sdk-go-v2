// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DisableCdpTask 操作DisableCdpTask
func (cli *ZSClient) DisableCdpTask(params param.DisableCdpTaskParam) (*view.DisableCdpTaskEventView, error) {
	resp := view.DisableCdpTaskEventView{}
	if err := cli.Post("v1/cdp-task/disable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

