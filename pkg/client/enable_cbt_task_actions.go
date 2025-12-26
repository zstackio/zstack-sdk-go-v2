// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// EnableCbtTask operates on EnableCbtTask
func (cli *ZSClient) EnableCbtTask(params param.EnableCbtTaskParam) (*view.EnableCbtTaskEventView, error) {
	resp := view.EnableCbtTaskEventView{}
	if err := cli.Post("v1/cbt-task/enable/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
