// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// EnableCbtTask 操作EnableCbtTask
func (cli *ZSClient) EnableCbtTask(params param.EnableCbtTaskParam) (*view.EnableCbtTaskEventView, error) {
	resp := view.EnableCbtTaskEventView{}
	if err := cli.Post("v1/cbt-task/enable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

