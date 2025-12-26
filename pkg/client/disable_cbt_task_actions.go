// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DisableCbtTask operates on DisableCbtTask
func (cli *ZSClient) DisableCbtTask(params param.DisableCbtTaskParam) (*view.DisableCbtTaskEventView, error) {
	resp := view.DisableCbtTaskEventView{}
	if err := cli.Post("v1/cbt-task/disable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
