// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DisableCbtTask 操作DisableCbtTask
func (cli *ZSClient) DisableCbtTask(params param.DisableCbtTaskParam) (*view.DisableCbtTaskEventView, error) {
	resp := view.DisableCbtTaskEventView{}
	if err := cli.Post("v1/cbt-task/disable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

