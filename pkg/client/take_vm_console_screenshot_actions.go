// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// TakeVmConsoleScreenshot 操作TakeVmConsoleScreenshot
func (cli *ZSClient) TakeVmConsoleScreenshot(uuid string, params param.TakeVmConsoleScreenshotParam) (*view.TakeVmConsoleScreenshotEventView, error) {
	resp := view.TakeVmConsoleScreenshotEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

